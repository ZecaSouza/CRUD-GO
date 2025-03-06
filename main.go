package main

import (
	"log"

	"github.com/ZecaSouza/CRUD-GO/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inicializa o router do Gin
	router := gin.Default()

	// Configura as rotas
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil{
		log.Fatal(err)
	}
}
