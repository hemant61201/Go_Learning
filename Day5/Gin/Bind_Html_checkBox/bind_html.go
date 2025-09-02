package main

import "github.com/gin-gonic/gin"

type form struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var f form

	c.Bind(&f)

	c.JSON(200, gin.H{
		"colors": f.Colors,
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("Day5/Gin/Bind_Html_checkBox/*")

	router.GET("/", indexHandler)

	router.POST("/", formHandler)

	router.Run(":8080")
}
