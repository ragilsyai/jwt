package middlewares

import (
	"github.com/gin-gonic/gin"
	"jwt/service"
	"net/http"
)

func Auhemtication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := service.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthentication",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
