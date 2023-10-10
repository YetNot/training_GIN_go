package main

import (
	"Gin/models"
	"Gin/routers"
	"Gin/storage"
	"log"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres", "host=... user=... password=... dbname=...")
	if err != nil {
		log.Println("error while accessing database", err)
	}
	defer storage.DB.Close()
	storage.DB.AutoMigrate(&models.Article{})

	router := routers.SetupRouter()
	router.Run()
}