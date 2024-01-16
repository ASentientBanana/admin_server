package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Validate(c *gin.Context) {

	auth_header := c.Request.Header.Get("Authorization")

	if auth_header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//parse jwt
	token, err := jwt.Parse(auth_header, func(token *jwt.Token) (interface {
	}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Extract the data
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Check if token expired
	// exp, err := strconv.Atoi(claims["exp"])

	// if err != nil {
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	// if time.Now().Unix() > int64(exp) {

	// }

	fmt.Println("Got:")
	fmt.Println(claims["sub"], claims["exp"])

	// db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	c.Next()
}
