package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseJson struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	DetailError map[string]interface{} `json:"detail_error"`
	Success bool `json:"success"`
}

type ResponseUnauthorized struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func SetResponse(c *gin.Context, code int, err error, message string, data interface{}) {
	if err != nil {
		c.JSON(code, ResultJson(nil, code, message, map[string]interface{}{"detail error": err.Error()}))
	} else {
		c.JSON(code, ResultJson(data, code, message, nil))
	}
}

func ResultJson(data interface{}, code int, message string, detailError map[string]interface{}) ResponseJson {
	var result ResponseJson
	result.Data = data
	result.Message = message
	result.DetailError = detailError
	result.Code = code
	if code != http.StatusOK {
		result.Success = false
	} else {
		result.Success = true
	}
	return result
}

