package repository

import (
	"github.com/piovani/go_api/domain"
)

type LivroRepository interface {
	Insert(livro *domain.Livro) (*domain.Livro, error)
}

type (repo LivroRepository) Insert(livro *domain.Livro) (*domain.Livro, error) {
	err := repo.Db.Create(livro).Error

	if err != nil {
		log.Fatalf("Error during persis livro: %v", err)
		return livro, err
	}

	return livro, nil
}