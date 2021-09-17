package main
import (
	"github.com/FirdavsMF/Installment-payments/pkg/billing"
	"fmt"
)
func main() {
	var amount float64
	var numbers float64
	// var percents int
    var monthe float64
	fmt.Print("Введите сумма товар: ")
    fmt.Scan(&amount) 
     
    fmt.Print("Введите номер телефон : ")
    fmt.Scan(&numbers)
    fmt.Print("Введите месяц ")
    fmt.Scan(&monthe)
	// fmt.Print("Введите процент ")
    // fmt.Scan(&percents)
    
	result := billing.Calculate2(amount, numbers, monthe)
    g := billing.Calculate1(amount)
    
	
	if monthe == 3 {
		fmt.Println("клиент должен оплатит:",result , "в месяц.","\nОбщая сумма составляет:", g)
	}

    result1 := billing.Calculate3(amount, numbers, monthe)
    g2 := billing.Calculate7(amount)

	if monthe == 6 {
        fmt.Println("клиент должен оплатит:",result1 , "в месяц.","\nОбщая сумма составляет:", g2)
	}
    
	result2 := billing.Calculate4(amount, numbers, monthe)
	g3 := billing.Calculate8(amount)
    if monthe == 9 {
		fmt.Println("клиент должен оплатит:",result2 , "в месяц.","\nОбщая сумма составляет:", g3)
	}
    
	result3 := billing.Calculate5(amount, numbers, monthe)
    g4 := billing.Calculate9(amount)
	if monthe == 12 {
		fmt.Println("клиент должен оплатит:",result3 , "в месяц.","\nОбщая сумма составляет:", g4)
	}
    
	result4 := billing.Calculate6(amount, numbers, monthe)
    g5 := billing.Calculate10(amount)
	if monthe == 18 {
		fmt.Println("клиент должен оплатит:",result4 , "в месяц.","\nОбщая сумма составляет:", g5)
	}
	
}