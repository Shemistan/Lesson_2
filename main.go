package main

import (
	"fmt"
	"strings"
)

type user struct {
	signed   bool
	log_pass string
}

func contains(s []user, elem string) bool {
	for _, logpass := range s {
		if logpass.log_pass == elem {
			return true
		}
	}
	return false
}

var products = []string{"apple", "banana", "grape"}
var getproducts string
var busket = make([]string, 0, cap(products))

func listProducts(prod []string) {
	var asd string
	for ind, prodElem := range prod {
		if ind < len(prod)-1 {
			asd += prodElem + ", "
		} else {
			asd += prodElem
		}
	}
	fmt.Println(asd)
}

func checkProduct(store []string, productToCheck string) bool {
	for _, item := range store {
		if item == productToCheck {
			return true
			break
		}
	}
	return false
}

func checkAllproducts(getProducts string, products []string) string {
	//busket = busket[:0]
	busket = strings.Split(getProducts, ",")
	answer := ""
	for _, item := range busket {
		isPresent := checkProduct(products, item)
		if !isPresent {
			answer += item
		}
	}
	return answer
}

func show() {
	listProducts(products)
	fmt.Println("--------------")
	fmt.Println("Please type the name of the product separated by comma.")
	fmt.Scan(&getproducts)
	checkedAnswer := checkAllproducts(getproducts, products)
	if checkedAnswer == "" {
		fmt.Println(busket, "added to busket")
	} else {
		fmt.Println("Please check product names. Wrong enries", checkedAnswer)
	}

	fmt.Println("Do you want to stay or exit the store. Type 2 to Exit or 3 to continue")
}

func main() {

	var command string = "begin"
	var cases int8
	var userList []user = make([]user, 0, 1000)

	for cases != 2 {
		//fmt.Println(command, "command")
		switch cases {
		case 0:
			fmt.Println("Welcome to Uzum store. \nPlease enter username and password separated by '_'.")
			fmt.Scan(&command)
			userobj := user{}
			userobj.signed = true
			userobj.log_pass = command
			userList = append(userList, userobj)
			fmt.Println("You are successfully signed up.\nPlease type 1 or 2  \n1. login\n2. Exit")
		case 1:
			fmt.Println("Please enter username and password separated by '_'.")
			fmt.Scan(&command)
			isThere := contains(userList, command)
			if isThere {
				fmt.Println("You are loged in successfully.\nWelcome to Uzum store")
				//fmt.Scan(&products)
				//fmt.Println(products, "have been added to your busket")
				//fmt.Println("Do you want to stay or exit the store\n3. Stay\n2. Exit")
				show()

			} else {
				fmt.Println("Username or password incorrect")
				continue
			}
		case 3:
			//fmt.Println("Please type the name of the product separated by comma\nApple\nBanana\nGrape")
			show()
		default:
			fmt.Println("enter valid")
		}
		fmt.Scan(&cases)
	}
}
