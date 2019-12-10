package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"net/http"
)

func handlePing(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func handleMainPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Code": "ok",
	})
}

func handleCrash(c *gin.Context) {
	fmt.Println("Crash")
	os.Exit(1)
	// panic("Oh Year I am Crashed")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handleMainPage)
	r.GET("/ping", handlePing)
	r.GET("/crash", handleCrash)

	r.Run(":8888")
}
