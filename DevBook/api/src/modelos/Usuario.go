package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuário representa um Usuário da aplicação
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}
	if err := usuario.formatar(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o campo Nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o campo Nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o campo Email é obrigatório e não pode estar em branco")
	}
	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("o campo Email está inválido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("o campo Senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
