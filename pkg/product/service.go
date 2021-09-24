package product

import (
	"errors"
	"log"
	"github.com/go-gomail/gomail"
)

type Product struct {
	Category	string			
	Price		int64					
}

func Calc(months int64, category string, price int64) (int64, error) {
	//product, err := s.ByID(ctx, id)
	//	if err != nil {
	//		log.Print(err)
	//	}
	products := Product{
		Category: category,
		Price: price,
	}
	 
	sum := int64(0)
	prodproc := int64(0)
	rangeMonths := int64(0)
	if (products.Category == "smartphone") { 
		prodproc = 3
		rangeMonths = 9
		sum = CategoryCalc(products, months, rangeMonths, prodproc)
		return sum, nil
	}
	if (products.Category == "computer") { 
		prodproc = 4
		rangeMonths = 12
		sum = CategoryCalc(products, months, rangeMonths, prodproc)
		return sum, nil
	}
	if (products.Category == "TV") {
		prodproc = 5
		rangeMonths = 18
		sum = CategoryCalc(products, months, rangeMonths, prodproc)
		return sum, nil
	}
	log.Println(0)
	return 0, errors.New("the request was executed incorrectly or this category is not in the list")
}

func CategoryCalc(product Product, months int64, rangeMonths int64, prodproc int64) (int64) {
		proc := int64(0)	
		proctime := int64(3)
		procplus := int64(0)
		if (months > rangeMonths) {
			if((months % 12 == 0) && (months > 12) || (months > 12)) {
				if ((months / 12 != 2) && (months > 24)) {
					proctime += 3 * (months / 12)
				}
				proctime += 3
				if ((months - rangeMonths != 6) && (months - rangeMonths != 12)) {
				procplus = 1
				}
			}
			procplus += (months - rangeMonths) / proctime			
			proc = procplus * prodproc
		}
		sumproc := product.Price * proc / 100
		sum := product.Price + sumproc
		return sum
}



func SendEmail(gmail, text string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "itsoflife@gmail.com")
	m.SetHeader("To", gmail)
	m.SetAddressHeader("Cc", gmail, "..")
	m.SetHeader("Subject", text)
	m.SetBody("text/html", text)
	m.Attach("myfavicon.ico")

	d := gomail.NewDialer("smtp.gmail.com", 587, "name", "password")

	
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return
}

// func SendEmail(message ,formAddress,receiveAddress ,username,authorizationCode string)  {
// 	// Create a new message object
// 	msg := gomail.NewMessage()

// 	//Sender
// 	msg.SetAddressHeader("From",formAddress,"False alarm")

// 	//recipient
// 	msg.SetHeader("To",msg.FormatAddress(receiveAddress,""))

// 	//text
// 	msg.SetBody("text/html",message)

// 	//Attachment msg.Attach

// 	//The default is QQ mailbox, you need to enter the mailbox "settings", find the "POP3/SMTP service" setting item in the "account" item, and turn it on
// 	d := gomail.NewDialer("smtp.qq.com", 465, username, authorizationCode) // Outgoing mail server, port, sender account number, authorization code
// 	if err := d.DialAndSend(msg); err != nil {
// 		log.Println("Failed to send", err)
// 		return
// 	}

// 	return

// }