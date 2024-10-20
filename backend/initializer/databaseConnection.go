package initializer

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseConnection() error {
	var dbName string = os.Getenv("DB_DATABSE")
	var dbPort string = os.Getenv("DB_PORT")
	var dbHost string = os.Getenv("DB_HOST")
	var dbPassword string = os.Getenv("DB_PASSWORD")
	var dbUser string = os.Getenv("DB_USER")
	var dsn string = constructDBUrl(dbName, dbPort, dbHost, dbPassword, dbUser)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
		return err
	}
	return nil
}

func constructDBUrl(dbName, dbPort, dbHost, dbPassword, dbUser string) string {
	return "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
}
