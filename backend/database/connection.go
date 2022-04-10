package db

import (
	"fmt"
	"github.com/LucaKuechler/StrengthTracker/models/DAO"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Client   *gorm.DB
	username = "root"
	password = "ordix"
	protocol = "tcp"
	host     = "localhost:3306"
	schema   = "ordix"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s)/%s", username, password, protocol, host, schema)

	client, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Client = client
	client.AutoMigrate(&dao.User{})

	log.Println("Database has been connected.")
}
