package main

import (
	"FIFA/handlers"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	emails := os.Getenv("TO")
	emailAddresses := strings.Split(emails, ",")

	for {
		err = handlers.SendMail(emailAddresses)
		if err != nil {
			log.Panic(err)
		}
	}
}
