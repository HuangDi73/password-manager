package main

import (
	"demo/password-manager/account"
	"demo/password-manager/files"
	"fmt"

	"github.com/fatih/color"
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
}

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDB("data.json"))
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
