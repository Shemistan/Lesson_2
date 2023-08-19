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
	currentUser := ""
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
			currentUser = ""
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			if userExists(userList, command) {
				fmt.Println("Этот пользователь уже существует")
			} else {
				userList = append(userList, command)

				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)

				fmt.Println(userList)
			}
		case auth:
			currentUser = ""
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			if userExists(userList, command) {
				currentUser = command
				fmt.Println("Добро пожаловать в магази")
			} else {
				fmt.Println("Вы не зарегистрированны")
			}

		case addProduct:
			if currentUser != "" {
				fmt.Println("Введите продукт")
				fmt.Scan(&command)
				productList = append(productList, command)
				fmt.Printf("%s добавлен в корзину\n", command)
			} else {
				fmt.Println("Вы не зарегистрированны")
			}
		case order:
			if currentUser != "" {
				if len(productList) == 0 {
					fmt.Println("Корзина пуста")
				} else {
					fmt.Printf("Чек (%s):\n", currentUser)
					for i := range productList {
						fmt.Println(i+1, productList[i])
					}
					productList = []string{}
				}
			} else {
				fmt.Println("Вы не зарегистрированны")
			}
		case cancel:
			if currentUser != "" {
				if len(productList) == 0 {
					fmt.Println("Корзина пуста")
				} else {
					productList = []string{}
					fmt.Println("Корзина очищена")
				}
			} else {
				fmt.Println("Вы не зарегистрированны")
			}
		}
	}
}

func userExists(users []string, newUser string) bool {
	for _, element := range users {
		if element == newUser {
			return true
		}
	}
	return false
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
