package migrations

import (
	"github.com/max10rogerio/book-go-example/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
