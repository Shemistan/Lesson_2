package main

import ("fmt")

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
	help = "help"
	addProduct = "add_product"
	show_products = "show_products"
	add_to_cart = "add_to_cart"
	cancel = "cancel"
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
		
		fmt.Println("Введите команду:")

		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case help:
			fmt.Println("команды:")
			fmt.Println("exit - Выход из программы")
			fmt.Println("auth - Авторизация")
			fmt.Println("reg - Регистрация")
			fmt.Println("add_product - Добавление продуктов")
			fmt.Println("show_products - Отображение продуктов")
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
			for ok := true; ok; ok = (input != "cancel"){
				
				fmt.Scan(&input)
				switch input {
					case cancel:
						fmt.Println(productList)
						break
					default:
						if check_if_exist(productList, input){
							productList = append(productList, input)
							fmt.Println("добавлен")
						}else {
							fmt.Println("продукт существует")
						}
							
				}
			}
		case show_products:
				fmt.Println(productList)
		case add_to_cart:
			fmt.Println("Напишите название продукта чтобы добавить в корзинку")
			var product string
			for ok := true; ok; ok = (product != "cancel"){
				fmt.Scan(&product)
				switch product {
				case cancel:
					fmt.Println(cartList)
					break
				default:
					if !check_if_exist(productList, product){
						if check_if_exist(cartList, product){
							cartList = append(cartList, product)
							fmt.Println("добавлен")
						}else {
							fmt.Println("продукт существует в корзинке")
						}
					}else{
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