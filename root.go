package main

import ( "net/http"
"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
router := gin.Default()
router.LoadHTMLGlob("templates/*")
	
router.GET("/a", func(c *gin.Context) {
c.HTML(http.StatusOK, "a.html", gin.H{ "title": "Page A", })
})
	
router.GET("/b", func(c *gin.Context) {
c.HTML(http.StatusOK, "b.html", gin.H{ "title": "Page B", })
})
	
router.GET("/c", func(c *gin.Context) {
c.HTML(http.StatusOK, "c.html", gin.H{ "title": "Page C", })
})
	
router.GET("/d", func(c *gin.Context) {
c.HTML(http.StatusOK, "d.html", gin.H{ "title": "Page D", })
})
	
router.Run(":8080")
}
