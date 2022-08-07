package repository

import (
	"log"

	"github.com/piovani/go_api/domain"
	"gorm.io/gorm"
)

type LivroRepository struct {
	Db *gorm.DB
}

func NewLicroRepository(db *gorm.DB) *LivroRepository {
	return &LivroRepository{
		Db: db,
	}
}

func (repo LivroRepository) Insert(livro *domain.Livro) (*domain.Livro, error) {
	err := repo.Db.Create(livro).Error

	if err != nil {
		log.Fatalf("Error during persis livro: %v", err)
		return livro, err
	}

	return livro, nil
}
