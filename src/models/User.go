package models

import (
	"api/src/auth"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (user *User) Prepare(state string) error {
	if erro := user.validate(state); erro != nil {
		return erro
	}

	if erro := user.format(state); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(state string) error {
	if user.Nome == "" {
		return errors.New("Nome é obrigatório")
	}

	if user.Nick == "" {
		return errors.New("Nick é obrigatório")
	}

	if user.Email == "" {
		return errors.New("Email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("E-mail inválido")
	}

	if state == "insert" && user.Senha == "" {
		return errors.New("Senha é obrigatório")
	}

	return nil
}

func (user *User) format(state string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if state == "insert" {
		hashedPassword, erro := auth.Hash(user.Senha)
		if erro != nil {
			return erro
		}

		user.Senha = string(hashedPassword)
	}

	return nil
}
