package main

import "fmt"

const (
	exit       = "exit"
	auth       = "auth"
	reg        = "reg"
	addProduct = "add_product"
	order      = "order"
)

func main() {
	var command string
	var checkAuth bool
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	_ = productList
outerLoop:
	for command != exit {
		fmt.Println("Введите команду")                                                                           // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Команда %s - выход из приложения\n", exit)                                                   // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Команда %s - авторизацияя\n", auth)                                                          // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Команда %s - регистрация\n", reg)                                                            // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Команда %s - добавление продукта (необходимо предваритально авторизоваться)\n", addProduct)  // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Printf("Команда %s - покупка и очистка корзины (необходимо предваритально авторизоваться)\n", order) // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует

			for _, v := range userList {
				if v == command {
					fmt.Println("Пользователь уже существует")
					continue outerLoop
				}
			}

			userList = append(userList, command)

			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			checkAuth = true
			fmt.Println(message)

			fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магази")
					checkAuth = true
					break
				} else {
					fmt.Println("Вы не зарегистрированны")
					checkAuth = false
					break
				}

			}
		case addProduct:
			if !checkAuth {
				fmt.Println("Вы не авторизованны")
				continue outerLoop
			}
			fmt.Println("Добавьте продует:")
			fmt.Scan(&command)
			for _, v := range productList {
				if v == command {
					fmt.Println("Такой продукт уже есть")
				}
			}

			productList = append(productList, command)

			message := fmt.Sprintf("Продукт %s успешно добавлен", command)
			fmt.Println(message)

			fmt.Println(productList)
		case order:
			if !checkAuth {
				fmt.Println("Вы не авторизованны")
				continue outerLoop
			}
			if len(productList) == 0 {
				fmt.Println("Корзина пуста")
				continue outerLoop
			}
			fmt.Printf("Вы купили %v\n", productList)
			productList = []string{}
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
