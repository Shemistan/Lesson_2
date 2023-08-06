package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
)

func dep() {
	var command string
	productList := make([]string, 0, 10)
	fmt.Println("Now please type the name of the products you'd like to buy.\nPlease enter them one by one, and the maximum amount of products is 10.\nIf you want to buy less than 10 products, after typing all your products type '0' to stop adding more products.\n")
	for command != "0" && len(productList) < 10 {
		fmt.Scan(&command)
		if command == "0" {
			break
		}
		productList = append(productList, command)
	}
	fmt.Println("Your products are: ")
	for _, i := range productList {
		fmt.Println(i)
	}
	fmt.Println("If you have finished choosing products, please type 'order' so that all your products will be ordered!")
	fmt.Scan(&command)
	if command == "order" {
		fmt.Println("Now your cart is free!")
		fmt.Println("You have ordered the following products: ")
		for _, i := range productList {
			fmt.Println(i)
		}
		fmt.Println("\nThank you for shopping with us!")
	}
	fmt.Println("Please type 'exit' to go to the main menu")
	fmt.Scan(&command)
}
func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2", "user_password"}

	for command != exit {
		fmt.Println("Enter a command:") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Println("Enter 'auth' to login into the system")
		fmt.Println("Enter 'reg' to register into the system")
		fmt.Println("Enter 'exit' to exit the system")
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)
			foundreg := false
			for _, v := range userList {
				if v == command {
					foundreg = true
					break
				}
			}

			if foundreg {
				fmt.Println("User already exists, please authenticate")
				break
			} else {
				userList = append(userList, command)
				message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
				fmt.Println(message)
				dep()
			} // Сделать так, что бы выводил сообщение, если пользователь уже существует

		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)
			found := false
			for _, v := range userList {
				if v == command {
					found = true
					break
				}
			}
			if found {
				fmt.Println("Welcome!")
				dep()
			} else {
				fmt.Println("Вы не зарегистрированны")
			}
		}

	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
