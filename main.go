package main

import "fmt"

const (
	exit = "exit"
	auth = "auth"
	reg  = "reg"
)

func main() {
	var command string
	userList := []string{"user1_password1", "user2_password2"}
	productList := make([]string, 0, 10)

	for command != exit {
		fmt.Println("Введите команду из следуюшего списка ")
		fmt.Println("reg -> для регистрации в системи ")
		fmt.Println("auth -> для авторизации в системи ")
		fmt.Println("exit -> для выхода из системи") // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			per1 := 0
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			for _, per := range userList {
				if per == command {
					fmt.Println("Ползовател с таким лог уже ест, попробуйте заново ...")
					fmt.Println()
					break
				} else {
					per1++
					if per1 == len(userList) {
						userList = append(userList, command)
						message := fmt.Sprintf("Пользователь %s успешно добавлен", command)
						fmt.Println(message)
					}
				}
			}

		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_passwor")
			fmt.Scan(&command)
			per2 := 0

			for _, v := range userList {
				bool1 := 0
				if v == command {
					for {
						if bool1 == 1 {
							break
						}
						fmt.Println("Добро пожаловать в магази")
						fmt.Println("Введите команду 'add_product' для добавление продукта")
						fmt.Println("Введите команду 'list' для просмотра добавленных продуктов")
						fmt.Println("Введите команду 'order' для заказа добавленных продуктов")
						var comand2 string
						fmt.Scan(&comand2)
						switch comand2 {
						case "add_product":
							fmt.Println("Введите продукт:")
							var per3 string
							fmt.Scan(&per3)
							productList = append(productList, per3)

						case "list":
							fmt.Println(productList)

						case "order":
							if len(productList) != 0 {
								fmt.Println("Проверте добавленные вами продуктов")
								fmt.Println(productList)
								fmt.Println("Все верно ?")
								fmt.Println("1 -> да")
								fmt.Println("2 -> нет")
								per4 := 0
								fmt.Scan(&per4)
								if per4 == 1 {
									fmt.Println("Ваш заказ оформлен")
									fmt.Println(productList)
									productList = nil
									//fmt.Println("PROVERKA etogo NIL")
									//fmt.Println(productList)
									bool1 = 1
								} else {
									fmt.Println("добавте продукт")
									continue
								}

							} else {
								fmt.Println("Вы не добавили продуктов в корзину, Пожалуста добавте!!!")
							}

						default:
							bool1 = 1
							fmt.Println("Не правилно ввели, попробуйте заново ...") // SHamistan podskajite pojalusta pochemu umenya ne srabativayet break zdes ???
							//bool1 = 1
							//break
						}
					}

				} else {
					per2++
					if per2 == len(userList) {
						fmt.Println("Вы не зарегистрированны")
					}
				}
				if bool1 == 1 {
					//fmt.Println("MI VIXODEM IZ tsikla ??????")
					break
				}
			}
		}
	}
}

// Реализовать следующие АПИ test
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
