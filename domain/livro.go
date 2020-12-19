package domain

type Livro struct {
	Domain
	Titulo string `json:"titulo" gorm:"varchar(255)"`
	Autor  string `json:"auto" gorm:"varchar(255)"`
}
