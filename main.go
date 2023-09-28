package main

import "fmt"

func main() {
	var name string
	var lastName string
	var number int
	var number1 int
	var exit string
	var yes string
	products := []string{"apple", "lemon", "bread", "meat", "milk"}
	cost := []int{200, 300, 400 ,500, 600}
	fmt.Println("Welcom to my Shop!!!")
	fmt.Println("Please enter your name!")
	fmt.Scan(&name)
	fmt.Println("Please enter your lastname!")
	fmt.Scan(&lastName)
	fmt.Println("You Are seller or customer")
	fmt.Println("1. Customer")
	fmt.Println("2. Seller")
	fmt.Scan(&number)
	if number == 1 {
		fmt.Println("List of products!")
		for i, val := range products{
			fmt.Println(i+1,".", val)
		}
		fmt.Println("List of costs product!")
		for i, value := range cost{
			fmt.Println(i+1, ".", value)
		} 
		productBox := make([]string, 0)
	fmt.Println("what you need!", "If you need any product, enter the product number!")
	fmt.Println("if yo don't any order product pleas write exit")
	fmt.Scan(&number1)
	if exit == "exit"{
		fmt.Println("goodbye!") 
	}else if number1 == 1{
		productBox = append(productBox, products[0])
	}else if number1 == 2 {
		productBox =append(productBox, products[1])
	}else if number1 == 3 {
		productBox = append(productBox, products[2])
	}else if number1 == 4 {
		productBox = append(productBox, products[3])
	}
	fmt.Println("if you want see your product box enter yes")
	fmt.Scan(&yes)
	if yes == "yes" {
		fmt.Println(productBox)
		fmt.Println("Bill",cost[number1])
	}
	}else if number == 2 {
		fmt.Println("Have good day colleague")
	}
}

// Реализовать следующие АПИ
// add_product - добавить продукт который вводиться с консоли в productList
// order - выводит сообщение что вы купили и очищает корзину
