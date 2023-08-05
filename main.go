package main

import (
	"fmt"
)

const (
	exit          = "exit"
	auth          = "auth"
	reg           = "reg"
	help          = "help"
	addProduct    = "add_product"
	show_products = "show_products"
	add_to_cart   = "add_to_cart"
	cancel        = "cancel"
	buy           = "buy"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)
	cartList := make([]string, 0, 5)
	fmt.Println(productList)
	_ = productList
	for command != exit {
		// Сделать красивый вывод, вывести список команд на этом шаге

		fmt.Print("Введите команду:")

		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case help:
			fmt.Println("команды:")
			fmt.Printf("%s - Выход из программы\n", exit)
			fmt.Printf("%s - Авторизация\n", auth)
			fmt.Printf("%s - Регистрация\n", reg)
			fmt.Printf("%s - Добавление продуктов\n", addProduct)
			fmt.Printf("%s - Отображение продуктов\n", show_products)
			fmt.Printf("%s - ДОбавление в корзинку\n", add_to_cart)
			fmt.Println()
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

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магази")

				} else {
					fmt.Println("Вы не зарегистрированны")
				}

			}
		case addProduct:
			fmt.Println("Добавьте продукты через [enter]")
			var input string
			var end string
			for ok := true; ok; ok = (input != "end") {

				fmt.Scan(&input)
				switch input {
				case end:
					fmt.Println(productList)
					break
				default:
					if check_if_exist(productList, input) {
						productList = append(productList, input)
						fmt.Println("добавлен")
					} else {
						fmt.Println("продукт существует")
					}

				}
			}
		case show_products:
			fmt.Println(productList)
		case add_to_cart:
			fmt.Println("Напишите название продукта чтобы добавить в корзинку")
			var product string
			for ok := true; ok; ok = (product != "buy") {
				fmt.Scan(&product)
				switch product {
				case buy:
					fmt.Println(cartList)
					fmt.Println("Заказ оформлен!")
					cartList = nil
					break
				default:
					if !check_if_exist(productList, product) {
						if check_if_exist(cartList, product) {
							cartList = append(cartList, product)
							productList = remove(productList, product)
							fmt.Println("добавлен")
						} else {
							fmt.Println("продукт существует в корзинке")
						}
					} else {
						fmt.Println("продукт не существует в списке")
					}
				}

			}

		}

	}
}

func check_if_exist(productList []string, product string) bool {
	for _, v := range productList {
		if v == product {
			return false
		}
	}
	return true
}

func remove(productList []string, product string) []string {
	for i, v := range productList {
		if v == product {
			productList = append(productList[:i], productList[i+1:]...)
			fmt.Println(productList)
			return productList
		}
	}
	return productList
}
