package main

import (
	"ticket-booking/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	route.Setup(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
