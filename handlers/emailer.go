package handlers

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

type EmailAgent struct {
	from     string
	password string
}

func SendMail(to []string) error {
	//defer time.Sleep(time.Hour * 3)		//rather just defer calling the function than make it sleep

	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	agent := EmailAgent{os.Getenv("FROM"), os.Getenv("PASSWORD")}
	auth := smtp.PlainAuth("", agent.from, agent.password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, agent.from, to, Message())
	if err != nil {
		return err
	}

	return nil
}

func Message() []byte {
	text, err := WebScraper()

	if err != nil {
		log.Panic(err)
	}

	subject := text
	body := ""
	message := []byte(subject + body)
	return message
}

func VerifyMessage() []byte {
	subject := "Please verify your email"
	body := "Please click the link below to verify your email\n"
	message := []byte(subject + body)
	return message
}
