package types

import (
	"encoding/json/v2"
	"time"
	"uuid"

	"github.com/moroz/kashanka/db/queries"
)

type AnswerMap map[queries.Locale]string

type CardDTO struct {
	ID               uuid.UUID
	QuestionLanguage queries.Locale
	AnswerLanguages  []queries.Locale
	Question         string
	Answers          AnswerMap
	InsertedAt       time.Time
	UpdatedAt        time.Time
}

func NewCardDTO(card *queries.Card) (*CardDTO, error) {
	dto := &CardDTO{
		ID:               card.ID,
		QuestionLanguage: card.QuestionLanguage,
		AnswerLanguages:  card.AnswerLanguages,
		Question:         card.Question,
		InsertedAt:       card.InsertedAt,
		UpdatedAt:        card.UpdatedAt,
	}

	if err := json.Unmarshal(card.Answers, &card.Answers); err != nil {
		return nil, err
	}

	return dto, nil
}
