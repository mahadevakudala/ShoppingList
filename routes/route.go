package routes

import (
	"shoppinglist/controller"

	"github.com/gin-gonic/gin"
)

var (
	signupController       = new(controller.SignupController)
	userController         = new(controller.UserController)
	shoppingListController = new(controller.ShoppingListController)
)

//InitRoutes give connection to database
func InitRoutes(r *gin.Engine) {
	UserGroup(r)
	ShoppingListGroup(r)

}
