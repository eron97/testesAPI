// routes/routes.go
package routes

import (
	"github.com/eron97/testesAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine, userController *controllers.UserController) {
	router.GET("/users", userController.GetUsersWithProcessing)
}
