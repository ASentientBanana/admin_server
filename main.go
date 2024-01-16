package main

import (
	"fmt"
	"os"

	"github.com/AsentientBanana/admin/server"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(os.Getenv("DEFAULT_ADMIN_USER"))
	server.InitDatabase()
	r := gin.Default()
	server.InitRoutes(r)
	r.Run(":9898")
}
