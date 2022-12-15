package password

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Hash []byte
}

func New(password string) (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return Password{}, nil
	}

	return Password{Hash: hash}, nil
}

func (p *Password) Compare(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}
