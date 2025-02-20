package account

import (
	"errors"
	"net/url"
	"time"
	"math/rand/v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ{}[]!@#$%^*()_-")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlStr string) (*Account, error) {
	if login == "" {
		return nil, errors.New("неверный формат логина")
	}
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}
	acc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlStr,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if login == "" {
		acc.generatePassword(12)
	}
	return acc, nil
}

func (acc *Account) generatePassword(passLen int) {
	pass := make([]rune, passLen)
	for i := range pass {
		pass[i] = letters[rand.IntN(len(letters))]
	}
	acc.Password = string(pass)
}
