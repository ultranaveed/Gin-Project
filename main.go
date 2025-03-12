package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	var address = ":3000"

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	log.Fatalln(router.Run(address))
}
