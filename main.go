package main

import (
	"log"
	"os"

	"github.com/AsentientBanana/admin/server"
)

func main() {
	log.SetOutput(os.Stdout)

	app := server.InitServer()

	panic(app.Run(":9898"))
}
