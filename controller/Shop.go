package controller

import (
	"fmt"
	"github.com/Shemistan/Lesson_2/model"
)

func AddProduct(product model.Product, db *DB) {
	if product.Name == "" || product.Category == "" {
		fmt.Println("Обязательные поля оставлять пустыми нельзя")
		return
	}

	if product.Amount == 0 {
		product.Amount = 1
	}

	for _, prod := range db.Products {
		if prod == product {
			prod.Amount += product.Amount
		}
		if prod.Name == product.Name && prod.Price != product.Price {
			fmt.Println("Данный продукт уже существует")
			return
		}
	}

	db.Products = append(db.Products, product)
}

func SellProduct(product model.Product, amount int, db *DB) model.Product {
	if product.Name == "" {
		fmt.Println("Неверный продукт")
		return model.Product{}
	}

	for _, prod := range db.Products {
		if prod.Name == product.Name && prod.Amount >= amount {
			prod.Amount -= amount
			return prod
		}
	}
	fmt.Println("Данного продукта не существует")

	return model.Product{}
}
