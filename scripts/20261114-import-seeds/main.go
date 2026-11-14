package main

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"uuid"

	"github.com/grailbio/base/tsv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moroz/kashanka/config"
)

type SeedRow struct {
	ID        string `tsv:"id"`
	Dutch     string `tsv:"dutch"`
	IPA       string `tsv:"ipa"`
	English   string `tsv:"en"`
	Chinese   string `tsv:"zh-Hant"`
	ExampleNL string `tsv:"example_nl"`
	ExampleEN string `tsv:"example_en"`
	ExampleZH string `tsv:"example_zh-Hant"`
	Level     string `tsv:"level"`
}

func gitRoot() string {
	var stdout bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(stdout.String())
}

func parseFile(filename string) ([]SeedRow, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []SeedRow

	reader := tsv.NewReader(file)
	reader.UseHeaderNames = true
	reader.HasHeaderRow = true

	for {
		var item SeedRow
		err := reader.Read(&item)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		result = append(result, item)
	}
	return result, nil
}

const insertCardQuery = `
insert into cards (id, question, answers)
	values ($1::uuid, $2, $3)
	on conflict (id) do update
	set question = excluded.question, answers = excluded.answers
	returning id
`

const insertCollectionQuery = `
insert into collections (title, slug)
	values ($1, $2)
	on conflict (slug) do update
	set title = excluded.title
	returning id
`

const insertJoinQuery = `
insert into collections_cards (collection_id, card_id, position)
	values ($1, $2, $3)
	on conflict (collection_id, position) do update
	set card_id = excluded.card_id
	returning id
`

func main() {
	root := gitRoot()
	ctx := context.Background()

	db, err := pgxpool.New(ctx, config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(ctx)

	files, err := filepath.Glob(filepath.Join(root, "seeds", "*.tsv"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		rows, err := parseFile(file)
		if err != nil {
			log.Fatal(err)
		}

		slug := strings.Replace(filepath.Base(file), ".tsv", "", 1)

		var collectionId uuid.UUID
		err = db.QueryRow(
			ctx,
			insertCollectionQuery,
			slug, slug,
		).Scan(&collectionId)
		if err != nil {
			log.Fatal(err)
		}

		for i, row := range rows {
			var cardId, joinTableId uuid.UUID
			answers, err := json.Marshal(map[string]string{
				"en":      row.English,
				"zh-Hant": row.Chinese,
			})
			if err != nil {
				log.Fatal(err)
			}

			err = db.QueryRow(ctx, insertCardQuery, row.ID, row.Dutch, answers).Scan(&cardId)
			if err != nil {
				log.Fatal(err)
			}

			err = db.QueryRow(ctx, insertJoinQuery, collectionId, row.ID, i+1).Scan(&joinTableId)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatal(err)
	}
}
