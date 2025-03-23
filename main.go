package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	var router = gin.Default()
	var address = ":3000"
	//c.Param
	//c,QUery
	//c.DefaultQuery
	//c.PostForm
	//c.DefaultPostForm
	//c.GetHeader
	//c.ShouldBindUri
	//c.ShouldBindJSON
	//c.ShouldBind
	//c.ShouldBindHeader

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Naveed")
	})

	router.GET("/products/:id", func(ctx *gin.Context) {
		var id = ctx.DefaultQuery("id", "100")
		fmt.Println("ID: ", id)
		ctx.String(http.StatusOK, "Got the ID!!")
	})

	router.Handle(http.MethodPost, "/post/:id", func(c *gin.Context) {
		var id = c.GetHeader("id")
		fmt.Println("ID :", id)
		c.String(http.StatusOK, "YEEE")
	})

	router.POST("/BProducts", func(ctx *gin.Context) {
		var product Product

		if e := ctx.ShouldBindJSON(&product); e != nil {
			ctx.String(http.StatusBadRequest, e.Error())
			return
		}

		fmt.Println("Product: ", product)
	})

	log.Fatalln(router.Run(address))
}
