package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestID = "request_id"

func RequestInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewString()
		c.Set(RequestID, requestID) // Gin handler 中使用
		ctx := context.WithValue(c.Request.Context(), RequestID, requestID)
		c.Request = c.Request.WithContext(ctx) // gRPC 调用链使用

		c.Next()
	}
}
