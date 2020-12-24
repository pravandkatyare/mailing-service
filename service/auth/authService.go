package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pravandkatyare/mailing-service/errs"
	"github.com/pravandkatyare/mailing-service/logging"
	"github.com/spf13/viper"
)

// ValidateHeader validates request with static header
func ValidateHeader(c *gin.Context) {
	if viper.GetString("header.api-key") != c.Request.Header.Get("x-api") {
		logging.Errorf("Aborted, unauthorized access attempt with api key: %s, %s", c.Request.Header.Get("x-api"), errs.ErrUnauthorized)
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized Access")
	}
	c.Next()
}
