package main

import (
	"log"
	"os"

	"github.com/eron97/testesAPI/config/controllers"
	"github.com/eron97/testesAPI/config/database"
	"github.com/eron97/testesAPI/config/routes"
	"github.com/eron97/testesAPI/config/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente:", err)
	}

	// Configurar o banco de dados
	dbConfig := database.Config{
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Database: os.Getenv("DATABASE"),
	}

	var db database.Database = &database.MySQLDB{} // Use a implementação real ou um mock

	_, err := db.InitDB(dbConfig)
	if err != nil {
		// Lidar com erro de inicialização do banco de dados
	}
	defer db.CloseDB()

	// Configurar o serviço e o controlador
	userService := &services.UserService{DB: db}
	userController := &controllers.UserController{Service: userService}

	// Configurar o roteador Gin
	router := gin.Default()
	routes.SetupTodoRoutes(router, userController) // Use a função de configuração do grupo de rotas

	// Iniciar o servidor
	router.Run(":8080")
}
