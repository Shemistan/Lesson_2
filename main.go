package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	add_product = "add_product"
	order = "order"
)

func main() {
	var command string
	var flag_auth bool //признак идентификации пользователя
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	//_ = productList
	for command != exit {
		fmt.Println("Введите команду, одну из exit/auth/reg/add_product/order") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			fmt.Println("exiting")
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			userList = append(userList, command)

			message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
			fmt.Println(message)

			fmt.Println("Привет")
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command {
					fmt.Println("Добро пожаловать в магазин")
					flag_auth=true			
				} 
			}
			if flag_auth!=true {
				fmt.Println("Вы не зарегистрированны или неверный пароль")
			}
		case add_product:
			//если авторизованы - добавляем товары
			if flag_auth==true {
				fmt.Println("Какой товар хотите добавить?")
				fmt.Scan(&command) 
				productList=append(productList, command)
					fmt.Println("Следующие товары в корзине",productList)
				} else {
					fmt.Println("Вы не авторизованы")
				}
		case order:
			//если авторизованы - добавляем товары
			if (flag_auth==true && len(productList)>0) {
				fmt.Println("Товары из корзины заказаны,", productList , "корзина очищена")
				productList = nil
				} else if flag_auth==false {
					fmt.Println("Вы не авторизованы")
				} else {
					fmt.Println("Нет товаров в корзине, добавьте товары через add_product")
				}

			}
	}
	
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
