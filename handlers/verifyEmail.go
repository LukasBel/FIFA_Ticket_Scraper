package handlers

import (
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

func VerifyEmail(to []string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	agent := EmailAgent{os.Getenv("FROM"), os.Getenv("PASSWORD")}
	auth := smtp.PlainAuth("", agent.from, agent.password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, agent.from, to, VerifyMessage())
	if err != nil {
		return err
	}

	return nil

}
