package initializers

import (
	"github.com/Onkar2104/go/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
