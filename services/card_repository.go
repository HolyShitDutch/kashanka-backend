package services

type CardRepository struct {
	db DB
}

func NewCardRepository(db DB) *CardRepository {
	return &CardRepository{
		db,
	}
}
