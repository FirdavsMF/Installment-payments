package main

import (
	"fmt"

	"github.com/FirdavsMF/Installment-payments/pkg/product"

)
func main() {
	var category string
	var price int64
	var months int64
	var gmail string
	fmt.Print("Введите категория товар (smartphone, computer, TV) : ")
    fmt.Scan(&category)
	fmt.Print("Введите сумма товар: ")
    fmt.Scan(&price) 
    fmt.Print("Введите диапазон ")
    fmt.Scan(&months)
	fmt.Print("Введите диапазон ")
    fmt.Scan(&gmail)
	text := ""
    t := product.SendEmail(gmail, text)
	// if err := product.SendEmail(gmail, text); err != nil {
	// 	panic(err)
	// }
    resault, err := product.Calc(months, category, price)
	fmt.Println(resault, err)
	
}

