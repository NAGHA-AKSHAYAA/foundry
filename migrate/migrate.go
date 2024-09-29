package migrate

import (
	"foundry/initialisers"
	"foundry/models"
)

func Migrate() {
	initialisers.DB.AutoMigrate(&models.User{})
	initialisers.DB.AutoMigrate(&models.Purchase{})
	initialisers.DB.AutoMigrate(&models.Stocks{})
}
