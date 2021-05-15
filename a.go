package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

const (
	domain = "http://oauth.think.com"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")
	codeUrl, _ := url.ParseRequestURI("http://127.0.0.1:8080/getcode")
	loginUrl := fmt.Sprintf("%s/auth?response_type=%s&client_id=%s&redirect_uri=%s", domain, "code", "clienta", codeUrl.String())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a-index.html", map[string]string{
			"loginUrl": loginUrl,
		})
	})

	r.GET("getcode", func(c *gin.Context) {
		code, _ := c.GetQuery("code")
		c.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	})

	r.Run(":8080")
}
