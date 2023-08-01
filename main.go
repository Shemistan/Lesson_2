package main

import (
	"fmt"
	"github.com/Shemistan/Lesson_2/controller"
	"github.com/Shemistan/Lesson_2/model"
	"os/exec"
)

const (
	Exit   = 0
	SignIn = 1
	SignUp = 2
	Sell   = 1
	Buy    = 2
	Cart   = 3
)

func cin[T any](msg string, variable T) {
	fmt.Print(msg)
	_, err := fmt.Scan(variable)
	if err != nil {
		fmt.Println("Ошибка ввода")
	}
}

func main() {
	var isAuth bool = false
	var db controller.DB
	var user model.User
	var cart []model.Product
	var command int = 99

	for command != Exit {
		exec.Command("clear").Run()
		if !isAuth {

			fmt.Println("Главное меню:\n[1] Войти\n[2] Зарегистрироваться\n[0] Выйти")
			cin("Введите команду: ", &command)

			switch command {
			case SignIn:
				{
					var logdata model.UserLogin

					cin("Введите логин: ", &logdata.Login)
					cin("Введите пароль: ", &logdata.Password)
					user = controller.SignIn(logdata, &db)
				}
			case SignUp:
				{
					var newUser model.User

					cin("Введите имя: ", &newUser.FullName)
					cin("Введите номер телефона: ", &newUser.PhoneNumber)
					cin("Введите баланс: ", &newUser.Balance)
					cin("Создайте логин: ", &newUser.Login)
					cin("Создайте пароль: ", &newUser.Password)

					user = controller.SignUp(newUser, &db)
				}
			case Exit:
				return
			default:
				fmt.Println("Неверная команда")
			}
			isAuth = true
			continue
		}

		if isAuth {
			fmt.Println("Меню:\n[1] Добавить продукт\n[2] Посмотреть список продуктов\n[3] Корзина\n[0] Выйти")
			cin("Введите команду: ", &command)

			switch command {
			case Sell:
				{
					var product model.Product

					cin("Введите название продукта: ", &product.Name)
					cin("Введите описание продукта: ", &product.Description)
					cin("Введите категорию продукта: ", &product.Category)
					cin("Введите количество продукта: ", &product.Amount)
					cin("Введите цену продукта: ", &product.Price)

					controller.AddProduct(product, &db)
				}
			case Buy:
				{
					var products []model.Product = db.Products
					var amount int
				shop:
					exec.Command("cls").Run()
					fmt.Println("Список продуктов:")

					for index, product := range products {
						fmt.Printf("[%v] %v\n", index+1, product)
					}
					fmt.Println("[0] Выйти")

					cin("Введите номер продукта: ", &command)
					if command > len(products) || command < 0 {
						fmt.Println("Неверный номер продукта")
						exec.Command("pause").Run()
						goto shop
					}

					if command == 0 {
						continue
					}

					if command > 0 && command <= len(products) {
						cin("Введите количество: ", &amount)
					}
					cart = append(cart, controller.SellProduct(products[command-1], command, &db))
				}
			case Cart:
				{
					var totalSum float64 = 0

				cart:
					exec.Command("cls").Run()

					fmt.Println("Корзина:")

					if len(cart) == 0 {
						fmt.Println("Корзина пуста")
						exec.Command("pause").Run()
						continue
					}

					for index, product := range cart {
						fmt.Printf("[%v] %v\n", index+1, product)
						totalSum += product.Price
					}
					fmt.Println("Общая сумма: ", totalSum)
					fmt.Printf("[%v] Оплатить\n", len(cart)+1)
					fmt.Println("[0] Выйти")

					fmt.Println("Если хотите убрать продукт из списка введите его номер, в ином случае можете оплатить или же выйти.")

					cin("Введите номер: ", &command)
					if command > len(cart)+1 || command < 0 {
						fmt.Println("Неверный номер продукта")
						exec.Command("pause").Run()
						goto cart
					}

					if command == 0 {
						continue
					}

					if command == len(cart)+1 {
						if totalSum > user.Balance {
							fmt.Println("Недостаточно средств")
							exec.Command("pause").Run()
							goto cart
						}

						user.Balance -= totalSum
						cart = cart[:0]
					}
				}
			case Exit:
				return
			default:
				fmt.Println("Неверная команда")
			}
		}
	}
}
