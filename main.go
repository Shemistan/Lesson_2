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
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("%v - для авторизации \n", auth)
		fmt.Printf("%v - для регистрации \n", reg)
		fmt.Printf("%v - для добавления продукта \n", add_product)
		fmt.Printf("%v - для покупки \n", order)
		fmt.Printf("%v - для выхода \n", exit)
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			isRegistred := false
			for _, v := range userList {
				if v == command {
					fmt.Println("Пользователь уже существует")
					isRegistred = true
					break
				}
			}
			if !isRegistred {
				userList = append(userList, command)

				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)

				fmt.Println(userList)
			}
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

			isAuth := false
			for _, v := range userList {
				if v == command {
					isAuth = true
					fmt.Println("Добро пожаловать в магазин")
				}
			}
			if !isAuth {
				fmt.Println("Вы не зарегистрированны")
			}
		case add_product:
			fmt.Println("Введите название продукта")
			fmt.Scan(&command)
			productList = append(productList, command)
			fmt.Printf("Продукт %s успешно добавлен \n", command)
			fmt.Println(productList)
		case order:
			fmt.Println("Вы купили", productList)
			productList = make([]string, 0, 10)
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
