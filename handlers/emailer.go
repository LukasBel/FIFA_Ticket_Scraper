package handlers

import (
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
	"time"
)

type EmailAgent struct {
	from     string
	password string
}

func SendMail(to []string) error {
	defer time.Sleep(time.Hour * 3)

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
