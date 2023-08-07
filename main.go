package main

import (
	"fmt"
)

const (
	exit        = "exit"
	auth        = "auth"
	reg         = "reg"
	add_product = "add_product"
	add_to_cart = "add_to_cart"
	order       = "order"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)
	orderList := make([]string, 0, 10)

	for command != exit {
		fmt.Println("Введите команду") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			message := ""
			if !string_in_list(command, userList) {
				userList = append(userList, command)
				message = fmt.Sprintf("Пользователь %s успешно добавлен", command)
			} else {
				message = fmt.Sprintf("Пользователь %s уже существует", command)
			}
			fmt.Println(message)

			fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)
			auth_flag := 0
			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магази")
					auth_flag := 1
					break
				}

				if auth_flag == 0 {
					fmt.Println("Вы не зарегистрированны")
				}
			}
		case add_product:
			fmt.Println("Введите название продукта")
			fmt.Scan(&command)
			message := ""
			if !string_in_list(command, productList) {
				productList = append(productList, command)
				message = fmt.Sprintf("Товар %s успешно добавлен в список товаров", command)
			} else {
				message = fmt.Sprintf("Товар %s уже существует", command)
			}
			fmt.Println(message)

			fmt.Println(productList)

		case add_to_cart:
			fmt.Println(productList)
			fmt.Println("Введите название продукта, который хотите выбрать")
			fmt.Scan(&command)
			message := ""
			if string_in_list(command, productList) {
				if !string_in_list(command, orderList) {
					orderList = append(orderList, command)
					message = fmt.Sprintf("Товар %s успешно добавлен в корзину", command)
				} else {
					message = fmt.Sprintf("Товар %s уже имеется в корзине", command)
				}
			} else {
				message = fmt.Sprintf("Товара %s нет в списке товаров", command)
			}
			fmt.Println(message)

			fmt.Println(orderList)

		case order:
			fmt.Println("У вас в корзине:")
			fmt.Println(orderList)
			fmt.Println("Нажмите 'yes' - для продолжения покупки, 'no' - для отмены")
			fmt.Scan(&command)
			for command != "yes" || command != "no" {
				if command == "yes" {
					fmt.Println("Спасибо за покупку")
					orderList = orderList[:0]
				}
				if command == "no" {
					fmt.Println("Спасибо за покупку")
				} else {
					fmt.Println("Вы ввели неправильную команду")
				}
			}

		}
	}
}

func string_in_list(get_str string, list []string) bool {
	for _, str := range list {
		if str == get_str {
			return true
		}
	}
	return false
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
// test comment
