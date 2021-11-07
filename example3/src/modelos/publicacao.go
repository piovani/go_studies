package modelos

import "time"

type Publicacao struct {
	ID        int64     `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   int64     `json:"autorId,omitempty"`
	AutorNick uint64    `json:"autorNick,omitempty"`
	Curtidas  int64     `json:"curtidas"`
	CriadoEm  time.Time `json:"criadaEm,omitempty"`
}
