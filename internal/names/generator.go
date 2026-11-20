package names

import (
	_ "embed"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

//go:embed adjectives.txt
var adjectivesRaw string
var adjectives = parseList(adjectivesRaw, "adjectives")

//go:embed nouns.txt
var nounsRaw string
var nouns = parseList(nounsRaw, "nouns")

func parseList(raw, desc string) []string {
	re := regexp.MustCompile("[\r\n]+")
	lines := re.Split(raw, -1)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	if len(result) == 0 {
		log.Fatalf("FATAL: The list of %s in the name generator module is empty!", desc)
	}
	return result
}

func sampleSlice(slice []string) string {
	index := rand.Int() % len(slice)
	return slice[index]
}

func GenerateName() string {
	return sampleSlice(adjectives) + " " + sampleSlice(nouns)
}
