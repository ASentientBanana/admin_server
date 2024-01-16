package controllers

import (
	"net/http"

	"github.com/AsentientBanana/admin/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem reading request",
		})
		return
	}

	var admin models.Admin

	// db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	// db.First(&admin, "Username = ?", body.Username)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"sub": admin.ID,
	// 	"exp": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// Sign and get the complete encoded token as a string using the secret
	// tokenString, err := token.SignedString(hmacSampleSecret)

	// fmt.Println(tokenString, err)

	// db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Sry brt bad request",
	// 	})
	// 	return
	// }

	// var admin models.Admin

	// db.First(&admin, "Username = ?", body.Username)

	// hashErr := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))
	// if hashErr != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Unauthorized, bad password" + " Supplied " + body.Password +" " +body.Username ,
	// 	})

	// 	return
	// }

	c.JSON(http.StatusOK, admin)
}
