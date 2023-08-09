package main

import (
	"fmt"
	"strings"
)

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	add_product = "add_product"
	order       = "order"
)

func main() {
	var command string
	var user string
	var hasValue = false
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	for command != exit {
		fmt.Println("\n *******  Menu   ******** \n")
		if hasValue {
			fmt.Println(add_product, "=====> добавить продукт в корзину")
			fmt.Println(order, "=====> показать список продуктов")
			fmt.Println(exit, "=====> для выхода программы")
			fmt.Print("\nВыбирите команду: ")
		} else {
			fmt.Println(auth, "=====> для авторизация")
			fmt.Println(reg, "========> для регистрация")
			fmt.Print("\nВыбирите команду: ")
		}
		fmt.Scan(&command)
		switch command {
		case exit:
			break
		case reg:
			fmt.Print("Введите логин и пароль в таком виде login_passwor: ")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					hasValue = true
				}
			}
			if hasValue {
				message := fmt.Sprintf("Пользователь с логином %s уже существует", command)
				fmt.Println("\n *****", message, "*****")
				fmt.Println("***** Попробуйте заново зарегистрироваться ***** \n")
			} else {
				userList = append(userList, command)
				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println("")
				fmt.Println("*****", message, "***** \n")
			}
		case auth:
			fmt.Print("Введите логин и пароль в таком виде login_passwor: ")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					hasValue = true
				}
			}

			if hasValue {
				message := fmt.Sprintf("\n ***** Добро пожаловать %s в магазин *****", command)
				fmt.Println("\n", message)
				user = command
			} else {
				fmt.Println("\n ***** Вы не зарегистрированны *****")
			}
		case add_product:
			if user == "" {
				fmt.Println("\n **** Пользователь не авторизован ****")
				break
			}
			fmt.Print("Введите название продукта: ")
			fmt.Scan(&command)

			productList = append(productList, command)
			message := fmt.Sprintf("Продукт %s успешно добавлен в корзину", command)
			fmt.Println("\n*****", message, "*****\n")
		case order:
			if user == "" {
				fmt.Println("\n **** Пользователь не авторизован ****")
				break
			}
			if len(productList) <= 0 {
				fmt.Println("В корзини ничего нет")
				break
			}
			products := fmt.Sprintf("%q", productList)
			fmt.Println("\nВы купили:", strings.Trim(fmt.Sprint(products), "[]"))
			productList = productList[:0]
			break
		default:
			fmt.Println("\n***** Вы не правильно ввели команду, ввидете правильную команду *****\n")
		}
	}
}
