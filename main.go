package main

import (
	"github.com/AsentientBanana/admin/server"
	"github.com/gin-gonic/gin"
)

func main() {
	server.InitDatabase()
	r := gin.Default()
	server.InitRoutes(r)
	r.Run(":9898")
}
