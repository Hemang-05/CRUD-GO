package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=mahmud.db.elephantsql.com user=asvdaolq password=FvRU7phjNq1uvv1jX6_wTEqBRwOoO4EZ dbname=asvdaolq port=5432 sslmode=disable " //os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in Connecting")
	}
}
