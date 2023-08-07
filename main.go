package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	help = "help"
	add_product = "add_product"
	order = "order"

)

func checkExist(List []string, item string) bool {
	for _, v := range List {
		if v == item {
			return true
		}
	}

	return false
}

func main() {
	var command string
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)

	_ = productList
	for command != exit {
		fmt.Println("Введите команду. Введите help для получения списка команд.")
		fmt.Scan(&command)

		switch command {
		case help:
			fmt.Println("\n",
				reg, "- Регистрация пользователя.\n",
				auth, "- Авторизация.\n",
				exit, "- Закончить работу.\n",
				add_product, "- Добавить продукт в корзину.\n",
				order, "- Заказать товары из корзины.\n")
		case reg:
			fmt.Println("Введите логин и пароль в формате login_password:")
			fmt.Scan(&command) 
			if checkExist(userList, command) == true {
				fmt.Println("Такой пользователь уже существует")
			} else {
				userList = append(userList, command)
				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)
			}
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			if checkExist(userList, command) == true {
				fmt.Println("Добро пожаловать в магазин")
			} else {
				fmt.Println("Вы не зарегистрированы")

			}
		case add_product:
			fmt.Println("Введите наименование товара:")
			fmt.Scan(&command)

			productList = append(productList, command)
			message := fmt.Sprintf("Товар %s успешно добавлен", command)
			fmt.Println(message)
		case order:
			if len(productList) == 0 {
				fmt.Println("Корзина пуста")
			} else {
				fmt.Println("Спасибо за заказ. Ваш список товаров:")
				for _, v := range productList {
					fmt.Println(v)
				}
			}
			productList = productList[:0]
		case exit:
			break
		default:
			fmt.Println("Введена неправильная команда.")
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
