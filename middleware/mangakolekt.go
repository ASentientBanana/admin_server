package middleware

import "github.com/gin-gonic/gin"

func ValidateOrigin(c *gin.Context) {

	c.Next()
}

func ValidateFiles(c *gin.Context) {

	c.Next()
}
