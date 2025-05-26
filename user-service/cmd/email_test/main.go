package main

import (
	"log"

	"user-service/internal/config"
	"user-service/internal/service"
)

func main() {
	cfg := config.NewConfig()
	emailService := service.NewEmailService(cfg)

	err := emailService.SendWelcomeEmail("bayadilov06@gmail.com", "TestUser")
	if err != nil {
		log.Fatal("Failed to send email:", err)
	}
	log.Println("Email sent successfully")
}
