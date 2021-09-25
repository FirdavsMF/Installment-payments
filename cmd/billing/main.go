package main

import (
	"fmt"

	"github.com/FirdavsMF/Installment-payments/pkg/product"

)
func main() {
	var category string
	var price int64
	var months int64
	// var gmail string
	fmt.Print("Введите категория товар (smartphone, computer, TV) : ")
    fmt.Scan(&category)
	fmt.Print("Введите сумма товар: ")
    fmt.Scan(&price) 
    fmt.Print("Введите диапазон ")
    fmt.Scan(&months)
	// fmt.Print("Введите электроний почта Gmail ")
    // fmt.Scan(&gmail)
    resault, err := product.Calc(months, category, price)
	// product.Gmail(category, months, price, gmail, resault )
    
	fmt.Println(resault, err)
	
}

