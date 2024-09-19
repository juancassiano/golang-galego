package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicação representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"id, omitempty"`
	Titulo    string    `json:"titulo, omitempty"`
	Conteudo  string    `json:"conteudo, omitempty"`
	AutorID   uint64    `json:"autorId, omitempty"`
	AutorNick string    `json:"autorNick, omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:criadoEm, omitempty`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.Validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil
}

// Validar é responsável por verificar se os campos da publicação estão preenchidos corretamente
func (publicacao *Publicacao) Validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

// formatar irá formatar os campos da publicação
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
