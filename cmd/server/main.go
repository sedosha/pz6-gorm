package main

import (
	"log"
	"net/http"
	"os"

	"example.com/pz6-gorm/internal/db"
	"example.com/pz6-gorm/internal/httpapi"
	"example.com/pz6-gorm/internal/models"
)

func main() {
	d := db.Connect()

	if err := d.AutoMigrate(&models.User{}, &models.Note{}, &models.Tag{}); err != nil {
		log.Fatal("migrate:", err)
	}

	r := httpapi.BuildRouter(d)

	// Получаем порт из переменной окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // fallback на 8081
	}

	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
