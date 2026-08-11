package response

import "github.com/gin-gonic/gin"

type Success struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"requestId"`
}

type Error struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

func SuccessWithData(gin *gin.Context, data interface{}, requestId string) {
	gin.JSON(200, Success{200, "success", data, requestId})
}

func SuccessWithoutData(gin *gin.Context, requestId string) {
	gin.JSON(200, Success{200, "success", nil, requestId})
}

func ErrorWithMessage(gin *gin.Context, message string, requestId string) {
	gin.JSON(500, Error{500, message, requestId})
}

func ErrorWithCode(gin *gin.Context, code int, requestId string) {
	gin.JSON(500, Error{code, "error", requestId})
}
