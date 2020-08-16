package middleware

import (
	"../shared"
	"github.com/gin-gonic/gin"
)

func Cors(config shared.CorConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.AllowOrigin)
		c.Header("Access-Control-Allow-Methods", config.AllowMethod)
		c.Header("Access-Control-Allow-Headers", config.AllowHeaders)
	}
}