package initializer

import (
	"backend/models"
	"log"
)

func SyncDatabase() error {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error syncing database : ", err)
		return err
	}
	return nil
}
