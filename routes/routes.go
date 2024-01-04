// routes/routes.go
package routes

import (
	"github.com/eron97/testesAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	// Grupo de rotas para usuários
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/:id", userController.GetUserByID)
		userGroup.PUT("/", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
		// Adicione outras rotas de usuário conforme necessário
	}

	// Adicione grupos de rotas para outros recursos conforme necessário

	return router
}
