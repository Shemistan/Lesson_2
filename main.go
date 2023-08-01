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
	var productCount uint
	var currentUserId int = -1
	userList := []string{"user1_password1"}
	productList := make([][]string, 10, 10)
	productCountList := make([][]uint, 10, 10)
	_ = productList
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			userList = append(userList, command)

			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)

			fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			for userId, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин")
					currentUserId = userId
				} else {
					fmt.Println("Вы не зарегистрированны")
				}
			}
		case add_product:
			if currentUserId == -1 {
				fmt.Println("Вы не авторизованы")
				return
			}

			fmt.Println("Введите название продукта и количество в таком виде Pizza 2")
			fmt.Scan(&command, &productCount)

			productList[currentUserId] = append(productList[currentUserId], command)
			productCountList[currentUserId] = append(productCountList[currentUserId], productCount)

		case order:
			if currentUserId == -1 {
				fmt.Println("Вы не авторизованы")
				return
			}

			fmt.Println("Заказ выполнен успешно. ваши продукты")
			fmt.Println("Название   -   Количество")
			for productId, product := range productList[currentUserId] {
				fmt.Println(product, " - ", productCountList[currentUserId][productId])
			}
		}

	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
