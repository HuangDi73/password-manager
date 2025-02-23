package main

import (
	"demo/password-manager/account"
	"demo/password-manager/encrypter"
	"demo/password-manager/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var variants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по url",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт по url",
	"5. Выход",
	"Введите вариант",
}

var actions = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountsByUrl,
	"3": findAccountsByLogin,
	"4": deleteAccounts,
}

func main() {
	fmt.Println("__Менеджер паролей__")
	if err := godotenv.Load(); err != nil {
		color.Red("Не удалость найти .env файл")
	}
	vault := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
Menu:
	for {
		choices := promptData(variants...)
		actionFunc := actions[choices]
		if actionFunc == nil {
			break Menu
		}
		actionFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	acc, err := account.NewAccount(login, password, url)
	if err != nil {
		color.Red("Неверный формат логина или url")
	}
	vault.AddAccount(*acc)
}

func deleteAccounts(vault *account.VaultWithDB) {
	url := promptData("Введите url для поиска")
	if isDeleted := vault.DeleteAccountsByUrl(url); isDeleted {
		color.Green("Аккаунты удалены")
	} else {
		color.Red("Аккаунты не найдены")
	}
}

func findAccountsByUrl(vault *account.VaultWithDB) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		return strings.Contains(acc.Url, url)
	})
	outputResult(&accounts)
}

func findAccountsByLogin(vault *account.VaultWithDB) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, login string) bool {
		return strings.Contains(acc.Login, login)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, acc := range *accounts {
		acc.Output()
	}
}

func promptData(prompts ...string) string {
	for i, p := range prompts {
		if len(prompts)-1 == i {
			fmt.Printf("%v: ", p)
		} else {
			fmt.Println(p)
		}
	}
	var input string
	fmt.Scanln(&input)
	return input
}
