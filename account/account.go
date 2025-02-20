package account

import (
	"errors"
	"net/url"
	"time"
)

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
	return acc, nil
}
