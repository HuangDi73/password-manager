package account

import "time"

type DB interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
