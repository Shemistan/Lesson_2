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
	userList := []string{"user1_password1", "user1_password1"}
	productList := make([]string, 0, 10)

	var auth_stat bool
	var reg_stat bool

	fmt.Println("    Добро пожаловать в Uzum-market!")

	for command != exit {
		fmt.Println(`Используйте следующие команды: "auth", "reg", "add_product", "order", "exit"`) // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break

		case reg: // Сделать так, что бы выводил сообщение, если пользователь уже существует
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)

			for _, v := range userList {
				if v == command && reg_stat == false {
					fmt.Println("Вы уже зарегистрированы! Пользователь уже существует.")
					reg_stat = true
				}
			}
			if reg_stat == false {
				userList = append(userList, command)
				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)
			}

			fmt.Printf("Зарегистрированные аккаунты:  %q \n", userList)

		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)
			for _, v := range userList {
				if v == command {
					fmt.Println("    Добро пожаловать в Uzum-market! Вы авторизовались!")
					auth_stat = true

				}
			}
			if auth_stat == false {
				fmt.Println("Вы не зарегистрированы!")
			}

		case add_product:
			fmt.Println("В корзине: ", productList)
			fmt.Println("Добавьте продукт в корзину: ")
			fmt.Scan(&command)
			productList = append(productList, command)
			fmt.Printf("Продукт - '%v' добавлен! ", command)
			fmt.Printf("%q \n", productList)

		case order:
			//fmt.Println("Поздравляю, вы купили: ")
			fmt.Printf("Поздравляю, вы купили: %q \n", productList)
			productList = productList[:0]
		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
