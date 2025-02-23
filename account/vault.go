package account

import (
	"demo/password-manager/encrypter"
	"encoding/json"
	"strings"
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
	db  DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) *VaultWithDB {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	decryptedData := enc.Decrypt(data)
	var vault Vault
	if err = json.Unmarshal(decryptedData, &vault); err != nil {
		color.Red("Не удалось разобрать vault файл")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDB) DeleteAccountsByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *VaultWithDB) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		isMatched := checker(acc, str)
		if isMatched {
			accounts = append(accounts, acc)
		}
	}
	return accounts
}

func (vault *Vault) toBytes() ([]byte, error) {
	data, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.toBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удаётся преобразовать в json")
	}
	vault.db.Write(encData)
}
