// controllers/user_controller.go
package controllers

import (
	"net/http"

	"github.com/eron97/testesAPI/config/services"
	"github.com/gin-gonic/gin"
)

// UserControllerInterface define a interface para o controlador de usuário
type UserControllerInterface interface {
	GetUsersWithProcessing(ctx *gin.Context)
	// Adicione outras funções de controlador conforme necessário
}

// UserController é a implementação real do controlador de usuário
type UserController struct {
	Service services.UserServiceInterface
}

// GetUsersWithProcessing obtém usuários com algum processamento adicional
func (c *UserController) GetUsersWithProcessing(ctx *gin.Context) {
	// Adapte para os requisitos específicos do seu aplicativo
	condition := ctx.Query("condition")
	users, err := c.Service.GetUsersWithProcessing(condition)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users with processing"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
