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
	userList := []string{"user1_password1", "user2_pass2"}
	productList := make([]string, 0, 10)

	fmt.Println("Список команд:")
	fmt.Printf("Команда %s - выход из приложения\n", exit)
	fmt.Printf("Команда %s - авторизацияя\n", auth)
	fmt.Printf("Команда %s - регистрация\n", reg)
	fmt.Printf("Команда %s - добавление продукта (необходимо предварительно авторизоваться)\n", addProduct)
	fmt.Printf("Команда %s - покупка и очистка корзины (необходимо предварительно авторизоваться)\n", order)
outerLoop:
	for command != exit {
		fmt.Println("Введите команду")
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)

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
					fmt.Println("Добро пожаловать в магазин")
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
					continue outerLoop
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
