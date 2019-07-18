package helper

import (
	"github.com/gin-gonic/gin"
	"../structs"
	"net/http"
)

func Responses(data structs.Response) gin.H{
		var (
			result gin.H
		)

		if data.Results != nil {
			result = gin.H {
				"error": false,
				"code": http.StatusOK,
				"message": data.Message,
				"total": len(data.Results),
				"result": data.Results,
			}
		} else {
			result = gin.H {
				"error": false,
				"code": http.StatusOK,
				"message": data.Message,
				"result": data.Result,
			}
		}
		return result
}

func ErrorCustomStatus(data structs.Response) gin.H {
	var (
		result gin.H
	)
	data.Error = true
	switch data.Code {
		case 404:
			data.Message = "Data tidak ditemukan"
		case 500:
			data.Message = "Internal Server Error"
		case 400:
			data.Message = "Bad Request"
		default:
			data.Message = "Terjadi kesalahan di internal server"
	}

	result = gin.H {
		"error": true,
		"code": data.Code,
		"message": data.Message,
	}

	return result
}

