package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	addProduct = "add_product"
	order = "order"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	for command != exit {
		fmt.Print("\n Введите команду:\n exit - Выход\n auth - Авторизация\n reg - Регистрация\n add_product - Добавить продукт в корзину\n order - Завершить покупку\n\n") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

sw:
		switch command {
		case exit:
			break
		case reg:
			fmt.Println("\n Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует

			for _, u := range userList {
				if u == command {
					fmt.Println("\n Пользователь уже существует!")
					break sw
				}
			}

			userList = append(userList, command)

			message := fmt.Sprintf("\n Пользователь %s успешно добавлен", command)
			fmt.Println(message)
			fmt.Println(userList)
		case auth:
			fmt.Println("\n Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("\n Добро пожаловать в магазин!")
				} else {
					fmt.Println("\n Вы не зарегистрированы!")
				}

			}
		case addProduct:
			fmt.Println("\n Введите название товара")
			fmt.Scan(&command)

			productList = append(productList, command)
		case order:
			fmt.Println("\n Список продуктов, которые вы приобрели:")

			for _, p := range productList {
				fmt.Println(p)
			}
			productList = []string{}
			fmt.Println("")
		default: 
			fmt.Println("\n Такой команды не существует!")
		}
	
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
