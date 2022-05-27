package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})


	router.GET("/page1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page1.html", gin.H{
			"title": "Page 1",
		})
	})


        router.GET("/syntax", func(c *gin.Context) {
		c.HTML(http.StatusOK, "syntax.html", gin.H{
			"title": "Syntax",
		})
	})


	router.GET("/comments", func(c *gin.Context) {
		c.HTML(http.StatusOK, "comments.html", gin.H{
			"title": "Comments In Go",
		})
	})

        router.GET("/variables", func(c *gin.Context) {
		c.HTML(http.StatusOK, "variables.html", gin.H{
			"title": "Variables",
		})
	})


	router.GET("/constants", func(c *gin.Context) {
		c.HTML(http.StatusOK, "constants.html", gin.H{
			"title": "Constants",
		})
	})


        router.GET("/datatypes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "datatypes.html", gin.H{
			"title": "Datatypes",
		})
	})


	router.GET("/arrays", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arrays.html", gin.H{
			"title": "Arrays In Go",
		})
	})


        router.GET("/slices", func(c *gin.Context) {
		c.HTML(http.StatusOK, "slices.html", gin.H{
			"title": "Slices",
		})
	})


	router.GET("/operators", func(c *gin.Context) {
		c.HTML(http.StatusOK, "operators.html", gin.H{
			"title": "Operators",
		})
	})


        router.GET("/conditions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "conditions.html", gin.H{
			"title": "Conditions",
		})
	})


	router.GET("/switch", func(c *gin.Context) {
		c.HTML(http.StatusOK, "switch.html", gin.H{
			"title": "Switch",
		})
	})

        router.GET("/loops", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loops.html", gin.H{
			"title": "Loops",
		})
	})


	router.GET("/functions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "functions.html", gin.H{
			"title": "Functions",
		})
	})


        router.GET("/struct", func(c *gin.Context) {
		c.HTML(http.StatusOK, "struct.html", gin.H{
			"title": "Structures in Go",
		})
	})


	router.GET("/maps", func(c *gin.Context) {
		c.HTML(http.StatusOK, "maps.html", gin.H{
			"title": "Maps",
		})
	})
	router.Run()
}