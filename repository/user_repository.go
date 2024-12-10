package repository

import (
	"fmt"
	"log"
	entity "trainig_api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email string) error
}

func Insert(db *gorm.DB) {
	user := entity.User{
		Email:    "",
		Password: "",
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("count:", result.RowsAffected)
}
