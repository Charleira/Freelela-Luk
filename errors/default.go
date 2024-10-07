package errors

import errors_model "github.com/Charleira/FreelelaLuk/model/errors-model"

var (
	BodyParser     = errors_model.ErrorModel{Id: 1, Msg: "Erro ao desserializar o body da requisição."}
	DecodeErrors   = errors_model.ErrorModel{Id: 2, Msg: "Erro ao decodificar os erros da resposta da api %s."}
	DecodeResponse = errors_model.ErrorModel{Id: 3, Msg: "Erro ao decodificar a resposta da api %s."}
	NotSystem      = errors_model.ErrorModel{Id: 4, Msg: "Sistema não configurado."}
	QueryParser    = errors_model.ErrorModel{Id: 5, Msg: "Erro o desserializar parametro via query"}
	Jwt            = errors_model.ErrorModel{Id: 6, Msg: "Erro ao ler o JWT."}
)
