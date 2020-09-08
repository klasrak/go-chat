package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/klasrak/go-chat/api/models"
)

var adminRole = models.Role{
	Role:        "admin",
	AccessLevel: 4,
}

var roles = []models.Role{
	{
		Role:        "admin",
		AccessLevel: 4,
	},
}

var users = []models.User{
	{
		Username: "admin",
		Password: "admin",
		RoleID:   1,
	},
}

// Load seeds initial data to DB
func Load(db *gorm.DB) {
	err := db.DropTableIfExists(&models.Role{}, &models.User{}).Error

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.AutoMigrate(&models.Role{}, &models.User{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range roles {
		err = db.Model(&models.Role{}).Create(&roles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed roles table: %v", err)
		}
	}

	for i := range users {
		user := &users[i]
		user.RoleID = 1
		err = db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatalf("cannot seed user table: %v", err)
		}
	}
}
