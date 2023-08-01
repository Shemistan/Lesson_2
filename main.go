package main

import (
	"fmt"
	"strings"
)

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	order       = "order"
	add_product = "add_product"
)

var menuInfo = map[string]string{
	"reg":         "регистрация нового профиля",
	"auth":        "авторизация в существующий профиль",
	"add_product": "добавить продукт в корзину",
	"order":       "создаёт чек и очищает корзину",
	"exit":        "выход из приложения",
}

func menuOutput() {
	helloMsgUpper := `|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
|          Добро пожаловать в Магазин        |`
	helloMsgLower := `|                           version by Faxa  |
|____________________________________________|`
	fmt.Printf("\033[1m%s\033[0m\n%s\n\n", helloMsgUpper, helloMsgLower)
	fmt.Println("╭─ \033[1mДоступные команды:\033[0m")
	for k, v := range menuInfo {
		// *: Команды будет выводиться рандомно т.к. это Map
		fmt.Printf("├─ \033[1m%s\033[0m - %s\n", k, v)
	}
	fmt.Print("Введите команду: ")
}

func main() {
	var command string
	var userData string
	var currentUser string
	userList := make(map[string]string)
	productList := make([]string, 0, 10)

	for command != exit {
		menuOutput()
		fmt.Scan(&command)

		switch command {
		case exit:
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&userData)

			if !strings.Contains(userData, "_") {
				fmt.Println("\n\x1b[31mНекорректный ввод. Ожидается формат 'login_password'\x1b[0m\n")
				break
			}

			loginPass := strings.Split(userData, "_")
			login := loginPass[0]
			password := loginPass[1]

			if _, ok := userList[login]; ok {
				fmt.Println("\n\x1b[33mЭтот пользователь уже существует\x1b[30m\n")
				break
			}
			userList[login] = password
			message := fmt.Sprintf("\n\x1b[32mПользователь %s успешно добавлен\x1b[0m", login)
			fmt.Println(message)

			fmt.Println(userList, "\n")
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&userData)

			if !strings.Contains(userData, "_") {
				fmt.Println("\n\x1b[31mНекорректный ввод. Ожидается формат 'login_password'\x1b[0m\n")
				break
			}

			loginPass := strings.Split(userData, "_")
			login := loginPass[0]
			fmt.Println()

			if _, ok := userList[login]; ok {
				fmt.Println("\n\x1b[32mДобро пожаловать в магазин\x1b[0m\n")
				currentUser = login
				break
			}
			fmt.Println("\x1b[33mВы ещё не зарегистрированы\x1b[0m\n")
		case add_product:
			if currentUser == "" {
				fmt.Println("\n\x1b[33mПользователь не авторизован\x1b[0m\n")
				break
			}
			fmt.Print("Введите продукт: ")
			var product string
			fmt.Scan(&product)

			productList = append(productList, product)
			fmt.Printf("\n\x1b[32m%s\x1b[0m был добавлен в корзину\n\n", product)
		case order:
			if currentUser == "" {
				fmt.Println("\n\x1b[33mПользователь не авторизован\x1b[0m\n")
				break
			}
			if len(productList) <= 0 {
				fmt.Println("\n\x1b[33mКорзина пуста\x1b[0m\n")
				break
			}
			fmt.Println()
			fmt.Printf("Чек пользователя - %s\n", currentUser)
			for i := range productList {
				fmt.Println("* " + productList[i])
			}
			fmt.Println()
			productList = []string{}
		default:
			fmt.Println("\n\x1b[31mТакой команды не существует\x1b[0m\n")
		}
	}
}