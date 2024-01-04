// main.go
package main

import (
	"github.com/eron97/testesAPI/controllers"
	"github.com/eron97/testesAPI/database"
	"github.com/eron97/testesAPI/routes"
	"github.com/eron97/testesAPI/services"
)

func main() {
	// Configurar o banco de dados
	dbConfig := database.Config{
		User:     "admin",
		Password: "adminadmin",
		Host:     "database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com",
		Port:     "3306",
		Database: "users",
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
	router := routes.SetupRouter(userController)

	// Iniciar o servidor
	router.Run(":8080")
}
