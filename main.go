package main

import (
	"fmt"
	"strings"
)

const (
	exit   = "exit"
	auth   = "auth"
	reg    = "reg"
	add    = "add"
	show   = "show"
	end    = "end"
	order  = "order"
	basket = "basket"
)

func main() {
	var command string
	listCommands := []string{
		exit,
		auth,
		reg,
		add,
		show,
		end,
		order,
		basket,
	}
	userList := []string{"user1_password1", "user1_password1"}
	regUserList := make(map[string]string)
	regUserList["undying"] = "123"
	productList := make([]string, 0, 10)
	productList = append(productList, "apple", "orange")
	testList := make([]string, len(productList))
	copy(testList, productList)
	currentUser := ""
	orderList := make(map[string][]string)
	orderList["undying"] = testList

	_ = productList
	for command != exit {
		fmt.Println("Введите команду", listCommands) // Сделать красивый вывод, вывести список команд на этом шаге
		fmt.Scan(&command)

		switch command {
		case exit:
			break
		case reg:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command) // Сделать так, что бы выводил сообщение, если пользователь уже существует
			namePassword := strings.Split(command, "_")
			name := namePassword[0]
			password := namePassword[1]
			regUserName, ok := regUserList[name]

			if ok && regUserName == password {
				fmt.Println("Пользователь уже существует")
			} else {
				userList = append(userList, command)
				regUserList[name] = password
				message := fmt.Sprintf("Пользователь %s успешно добавлен", name)
				fmt.Println(message)
			}

			// fmt.Println(userList)
		case auth:
			fmt.Println("Введите логин и пароль в таком виде login_password")
			fmt.Scan(&command)
			namePassword := strings.Split(command, "_")
			name := namePassword[0]
			password := namePassword[1]
			regUserName, ok := regUserList[name]
			if ok && regUserName == password {
				fmt.Println("Добро пожаловать в магазин")
				currentUser = name
			} else {
				fmt.Println("Вы не зарегистрированы")
			}
		case add:
			if currentUser != "" {
				fmt.Println("Введите название нового продукта")
				fmt.Scan(&command)
				product := command
				ok := false
				for _, v := range productList {
					if v == product {
						fmt.Println("Этот продукт уже существует")
						ok = true
						break
					}
				}
				if !ok {
					fmt.Println("Продукт успешно добавлен")
					productList = append(productList, product)
				}
			} else {
				fmt.Println("Если хотите добавить новый продукт, сначала пройдите авторизацию")
			}
		case order:
			if currentUser != "" {
				for {
					fmt.Println("Введите то, что хотите купить из списка продуктов или введите end, чтобы завершить заказ", productList)

					fmt.Scan(&command)
					if command == "end" {
						break
					}
					ok := true
					for _, v := range productList {
						if command == v {
							orderList[currentUser] = append(orderList[currentUser], command)
							fmt.Println("Продукт успешно добавлен в вашу корзину")
							ok = false
						}
					}
					if ok {
						fmt.Println("Этот продукт не существует в списке продуктов")
					}
				}

			} else {
				fmt.Println("Если хотите что-то заказать, сначала пройдите авторизацию")
			}
		case basket:
			if currentUser != "" {
				for {
					fmt.Println("Введите то, что хотите удалить из корзинка или all (чтобы удалить всех продуктов) или введите end, чтобы завершить")
					fmt.Println(orderList[currentUser])
					fmt.Scan(&command)

					if command == "end" {
						break
					}
					if command == "all" {
						orderList[currentUser] = append(orderList[currentUser][:0], "")
						break
					}
					for i, v := range orderList[currentUser] {
						if command == v {
							orderList[currentUser] = append(orderList[currentUser][:i], orderList[currentUser][i+1:]...) // in progress
							fmt.Printf("%v удалено из корзины\n", v)
						}
					}
					fmt.Println(orderList[currentUser])

					if len(orderList[currentUser]) == 0 {
						fmt.Println("Корзинка пустая")
						break
					}
				}
			} else {
				fmt.Println("Если хотите увидеть товары в вашей корзине, сначала пройдите авторизацию")
			}
		case end:
			if currentUser != "" {
				fmt.Println("Вы успешно вышли из аккаунта")
				currentUser = ""
			} else {
				fmt.Println("Если вы хотите выйти из своей аккаунта, вы сначала войдете")
			}

		case show:
			fmt.Println("Что хотите видеть")
			fmt.Scan(&command)
			switch command {
			case "product":
				fmt.Println("productList: ", productList)
			case "user":
				fmt.Println("RegUserList: ", regUserList)
			case "current":
				fmt.Println("CurrnetUser: ", currentUser)
			case "order":
				fmt.Println("OrderList: ", currentUser, orderList[currentUser])
			}

		}
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
