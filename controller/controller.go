package controller

import (
	"shoppinglist/model/wrapper"

	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	userWrapper *wrapper.UserWrapper
)

//InitializeController connects the database to shoppinglist
func InitializeController(DB *gorm.DB) {
	db = DB
	userWrapper = wrapper.CreateUserWrapper(db)
}
