package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (int64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUE (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return int64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID int64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		SELECT p.id, p.titulo, p.conteudo, p.autor_id, p.curtidas, u.nick FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		WHERE p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}
