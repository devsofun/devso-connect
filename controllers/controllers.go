package controllers

import (
	"log"

	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name":    "Devso Connect Platform",
		"Version": "0.0.1",
	})
}

func ShowLoginForm(c *gin.Context) {
	checkStr := make([]byte, 20)
	for i := 0; i < 20; i++ {
		checkStr[i] = byte('a' + rand.Intn(26))
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"checkStr": string(checkStr),
	})
}

func ProcessLogin(c *gin.Context) {
	pasteUrl := c.PostForm("pasteUrl")
	checkStr := c.PostForm("checkStr")
	log.Println("[Luogu Login] Pastebin URL: ", pasteUrl, "Token: ", checkStr)
}
