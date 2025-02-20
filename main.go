package main

import "fmt"

var variants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по url",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт по url",
	"5. Выход",
	"Введите вариант",
}

func main() {
	choices := promptData(variants...)
	fmt.Println(choices)
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
