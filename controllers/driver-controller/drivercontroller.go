package driverController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Cria um novo usuário
// @Description	Cria um novo usuário no sistema com as informações fornecidas
// @Accept		multipart/form-data
// @Param		systemid		path		string							true	"Id do Sistema"
// @Param		files			formData	file							true	"Arquivos de usuários para upload"
// @Success		201				{object}	int
// @Failure		400				{object}	[]error_model.ErrorModel
// @Router		/system/{systemid}/users [post]
func CreateUser(c *gin.Context) {

	// Verifica se foram fornecidos arquivos
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar os arquivos enviados"})
		return
	}

	// Inicializa um slice para armazenar os resultados
	var uploadResults []interface{}

	// Itera sobre os arquivos enviados
	for _, fileHeaders := range form.File["files"] {
		file, err := fileHeaders.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir o arquivo"})
			return
		}
		defer file.Close()

	}

	// Retorna a resposta com os resultados do upload
	c.JSON(http.StatusOK, uploadResults)
}
