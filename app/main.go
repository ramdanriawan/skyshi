package main

import (

	config "skyshi.com/src/config"
	route "skyshi.com/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.DB()

	route.Api(r, db)

	r.Run(":3030")
}
