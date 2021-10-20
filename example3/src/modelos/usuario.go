package modelos

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:""nome,omitempty`
	Nick     string    `json:""nike,omitempty`
	Email    string    `json:""email,omitempty`
	Senha    string    `json:""senha,omitempty`
	CriadoEm time.Time `json:""criadoEm,omitempty`
}
