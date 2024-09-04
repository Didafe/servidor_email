package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	from := "fernando@localhost.com"
	to := []string{"recebedor@localhost.com"}

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to[0] + "\r\n" +
			"Subject: Teste de envio de email\r\n" +
			"\r\n" +
			"Mensagem de teste",
	)

	auth := smtp.PlainAuth("", "", "", "localhost")
	err := smtp.SendMail("localhost:1025", auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email enviado com sucesso")
}