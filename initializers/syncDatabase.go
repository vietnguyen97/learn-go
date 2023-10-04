package initializers

import (
	"github.com/nguyen997/gin-gorm-rest/config"
	"github.com/nguyen997/gin-gorm-rest/models"
)

func SyncDatabase() {
	config.DB.AutoMigrate(&models.User{})
}
