package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pravandkatyare/mailing-service/service/auth"
)

// AuthMiddleware for handling header authentication of static keyword
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.ValidateHeader(c)
		c.Next()
	}
}
