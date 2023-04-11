package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func main() {
	user := ""
	password := ""
	msg := gomail.NewMessage()

	msg.SetHeader("From", "from@example.com")
	msg.SetHeader("To", "to@example.com")
	msg.SetHeader("Subject", "Assunto Teste")
	msg.SetBody("text/html", "<b>Corpo teste</b>")

	n := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, user, password)

	if err := n.DialAndSend(msg); err != nil {
		fmt.Printf("DEU RUIM: %v", err)
	} else {
		fmt.Println("DEU BOM!!!")
	}
}
