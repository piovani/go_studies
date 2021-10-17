package domain

import (
	"errors"
	"log"
)

type Livro struct {
	Domain
	Titulo string `json:"titulo" gorm:"varchar(255)"`
	Autor  string `json:"auto" gorm:"varchar(255)"`
}

func (livro *Livro) validate() error {
	if livro.Autor == "" {
		log.Fatalf("Error during validation Livro Autor")
		return errors.New("Error during validation Livro Autor")
	}

	if livro.Titulo == "" {
		log.Fatalf("Error during validation Livro Titulo")
		return errors.New("Error during validation Livro Titulo")
	}

	return nil
}
