package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Charleira/FreelelaLuk/controllers"
	"github.com/Charleira/FreelelaLuk/docs"

	"github.com/gin-gonic/gin"
)

var PORT = 3018
var gAppVersion string = "development"

func main() {
	// Verifica se a variável de ambiente APP_PORT foi definida
	if port, ok := os.LookupEnv("APP_PORT"); ok {
		PORT, _ = strconv.Atoi(port)
	}

	// Configuração do Swagger
	if _, ok := os.LookupEnv("USE_DEBUG"); ok {
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", PORT)
	} else {
		docs.SwaggerInfo.BasePath = "/api"
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "myapp-production.com"
	}

	// Inicializa o servidor Gin
	router := gin.Default()

	// Adiciona as rotas dos controladores
	controllers.AddRoutes(router)

	// Exibe informações no console
	fmt.Printf("API Version: %s\n", gAppVersion)
	fmt.Printf("Rodando na porta: %d\n", PORT)

	// Inicia o servidor na porta configurada
	router.Run(fmt.Sprintf(":%d", PORT))
}
