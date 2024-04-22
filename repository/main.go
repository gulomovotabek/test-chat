package repository

import (
	"cae_chat/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	// Add your methods here
}

func NewRepository() Repository {
	db, err := New()
	if err != nil {
		panic(err)
	}

	return &repository{db}
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) CreateMessage(msg *models.Message) error {
	return nil
}
