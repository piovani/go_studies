package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        int64     `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   int64     `json:"autorId,omitempty"`
	AutorNick uint64    `json:"autorNick,omitempty"`
	Curtidas  int64     `json:"curtidas"`
	CriadoEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O Titulo é obrigatório")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O Conteudo é obrigatório")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
