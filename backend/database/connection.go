package db

import (
	"fmt"
	"github.com/LucaKuechler/StrengthTracker/models/DAO"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var Client *gorm.DB

func init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	schema := os.Getenv("DB_NAME")
    protocol := "tcp"

	dataSourceName := fmt.Sprintf("%s:%s@%s(%s)/%s", username, password, protocol, host, schema)

	client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Client = client
	client.AutoMigrate(&dao.User{})

	log.Println("Database has been connected.")
}
