package domain

type Livro struct {
	Domain
	Titulo string `json:"titulo" gorm:"varchar(255)"`
	Autor  string `json:"auto" gorm:"varchar(255)"`
}

func NewLivro(titulo string, autor string) *Livro {
	return &Livro{
		Titulo: titulo,
		Autor:  autor,
	}
}

func validate() error {
	return nil
}
