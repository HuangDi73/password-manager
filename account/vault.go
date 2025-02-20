package account

import (
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type DB interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db DB
}

func NewVault(db DB) *VaultWithDB {
	data, err := db.Read()
	if err != nil {
		color.Red("Не удалось прочитать файл")
		return &VaultWithDB{
			Vault: &Vault{
				Accounts:  []Accounts{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db: DB,
		}
	}
	var vault Vault
	err := json.Unmarshal(data, &vault)
	if err != nil {
		color.Red("Не удалось разобрать json")
		return &VaultWithDB{
			Vault: &Vault{
				Accounts:  []Accounts{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db: DB,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    DB,
	}
}

func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) toBytes() ([]byte, error) {
	data, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vault *VaultWithDB) save() {
	vault.updatedAt = time.Now()
	data, err := vault.toBytes()
	if err != nil {
		color.Red("Не удаётся преобразовать в json")
	}
	vault.db.Write(data)
}
