package routes

import (
	"shoppinglist/controller"

	"github.com/gin-gonic/gin"
)

var (
	signupController = new(controller.SignupController)
)

//InitRoutes give connection to database
func InitRoutes(r *gin.Engine) {
	UserGroup(r)

}
