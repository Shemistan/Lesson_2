package main

import "fmt"

const (
	exit       = "exit"
	auth       = "auth"
	reg        = "reg"
	addProduct = "add_product"
	order      = "order"
	cancel     = "cancel"
)
const (
	regDescription        = reg + " - регистрация нового профиля"
	authDescription       = auth + " - авторизация в существующий профиль"
	addProductDescription = addProduct + " - добавить продукт в корзину"
	orderDescription      = order + " - создаёт чек и очищает корзину"
	exitDescription       = exit + " - выход из приложения"
	cancelDescription     = cancel + " - очищает корзину"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Println(regDescription)
		fmt.Println(authDescription)
		fmt.Println(addProductDescription)
		fmt.Println(orderDescription)
		fmt.Println(exitDescription)
		fmt.Println(cancelDescription)
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
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магази")

				} else {
					fmt.Println("Вы не зарегистрированны")
				}

			}
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
