package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	help = "help"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Println("Список доступных команд: ")
		fmt.Println("reg - register a new user")
		fmt.Println("auth - login with an existing user")
		fmt.Println("add_product - add a new product")
		fmt.Println("order - order products")
		fmt.Println("exit - exit the program")
		fmt.Println("help - see available commands")

		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			if checkIfUserExists(userList, command) {
				fmt.Println("Пользователь с таким именем уже существует")
				break
			}
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			userList = append(userList, command)

			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)

			fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

			if checkIfUserExists(userList, command) {
				fmt.Println("Добро пожаловать в магазин!")
			} else {
				fmt.Println("Пользователь не найден")
			}
		case "add_product":
			fmt.Println("Пожалуйста введите название продукта: ")
			fmt.Scanln(&command)
			productList = append(productList, command)
			fmt.Printf("Продукт %s был добавлен успешно! \n", command)
		case "order":
			fmt.Println("Вы успешно совершили покупку!")
			fmt.Println("Корзина: ")
			for i := 0; i < len(productList); i++ {
				if productList[i] != "" {
					fmt.Println(productList[i])
				}
			}
			productList = make([]string, 2)
		}
	}
}

func checkIfUserExists(userList []string, username string) bool {
	for _, v := range userList {
		if v == username {
			return true
		}
	}
	return false
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
