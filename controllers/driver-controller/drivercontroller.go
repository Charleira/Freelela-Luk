package driverController

import (
	"net/http"

	"github.com/Charleira/FreelelaLuk/errors"
	errors_model "github.com/Charleira/FreelelaLuk/model/errors-model"
	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Cria um novo usuário
// @Description	Cria um novo usuário no sistema com as informações fornecidas
// @Accept		multipart/form-data
// @Param		systemid		path		string							true	"Id do Sistema"
// @Param		files			formData	file							true	"Arquivos de usuários para upload"
// @Success		201				{object}	int
// @Failure		400				{object}	[]errors_model.ErrorModel
// @Router		/users [post]
func CreateUser(c *gin.Context) {

	// Verifica se foram fornecidos arquivos
	form, err := c.MultipartForm()
	if err != nil {

		c.JSON(http.StatusBadRequest, []errors_model.ErrorModel{errors.BodyParser})
		return
	}

	// Itera sobre os arquivos enviados
	for _, fileHeaders := range form.File["files"] {
		file, err := fileHeaders.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, []errors_model.ErrorModel{errors.BodyParser})
			return
		}
		defer file.Close()

	}

	// Retorna a resposta com os resultados do upload
	c.JSON(http.StatusOK, 1)
}
