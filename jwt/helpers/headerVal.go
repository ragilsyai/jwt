package helpers

import "github.com/gin-gonic/gin"

func GetContenrType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
