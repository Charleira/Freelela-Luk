package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Charleira/FreelelaLuk/controllers"
	"github.com/Charleira/FreelelaLuk/docs"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
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

	// Inicializa o servidor Fiber
	app := fiber.New()

	// Configura o Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Adiciona as rotas dos controladores
	controllers.AddRoutes(app)

	// Exibe informações no console
	fmt.Printf("API Version: %s\n", gAppVersion)
	fmt.Printf("Rodando na porta: %d\n", PORT)

	// Inicia o servidor na porta configurada
	err := app.Listen(fmt.Sprintf(":%d", PORT))
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
