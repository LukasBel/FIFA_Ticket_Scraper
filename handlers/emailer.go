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

func SendMail(to []string, message []byte) error {
	//defer time.Sleep(time.Hour * 3)		//rather just defer calling the function than make it sleep

	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	agent := EmailAgent{os.Getenv("FROM"), os.Getenv("PASSWORD")}
	auth := smtp.PlainAuth("", agent.from, agent.password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, agent.from, to, message)
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

func WelcomeMessage() []byte {
	subject := "Welcome to the FIFA newsletter!"
	body := "Thank you for subscribing to the FIFA newsletter! You will now receive the latest news and updates about the 2026 FIFA World Cup."
	message := []byte(subject + body)
	return message
}
