// controllers/user_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/eron97/testesAPI/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

// CreateUser cria um novo usuário
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user services.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// GetUserByID obtém um usuário pelo ID
func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.Service.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUser atualiza as informações de um usuário
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user services.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser exclui um usuário pelo ID
func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.Service.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// ... outras funções de controlador
