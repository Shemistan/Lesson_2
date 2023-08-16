package main

import "fmt"

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	add_product = "add_product"
	order       = "order"
)

func main() {
	var command string
	var command1 string
	userList := []string{"user1_password1"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду! \n Для авторизации - 'auth' \n Для регистрации - 'reg' \n Для выхода - 'exit' ")
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			for _, user := range userList {
				if user == command {
					fmt.Println("Пользователь уже существует")
					return
				}
			}
			userList = append(userList, command)

			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)

		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин!")
					fmt.Println("Введите команду \n Для добавление продукта - add_product \n Для вывода корзины - order")
					switch command1 {
					case add_product:
						fmt.Println("Введите продукт")
						fmt.Scan(&command1)
						productList = append(productList, command1)
						message := fmt.Sprintf("Продукт %s успешно добавлен", command1)
						fmt.Println(message)

					case order:
						message := fmt.Sprintf("Вы купили %s", productList)
						fmt.Println(message)
						fmt.Println("Корзина очищена")
						fmt.Println(productList)
						break
					}
				} else {
					fmt.Println("Вы не зарегистрированны")
					break
				}

			}
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
