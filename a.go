package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	authServerUrl = "http://oauth.think.com" // 认证中心地址
	stateA        = "myclient"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     "clienta",
		ClientSecret: "123",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerUrl + "/auth",  // 获取授权码
			TokenURL: authServerUrl + "/token", // 获取token
		},
		RedirectURL: "http://127.0.0.1:8080/getcode",
		Scopes:      []string{"all"},
	}
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")
	// 构造登陆地址
	loginUrl := oauth2Config.AuthCodeURL(stateA)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a-index.html", map[string]string{
			"loginUrl": loginUrl,
		})
	})

	r.GET("getcode", func(c *gin.Context) {
		if state, ok := c.GetQuery("state"); !ok || state != stateA {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "wrong state",
			})
		}
		code, _ := c.GetQuery("code")
		token, err := oauth2Config.Exchange(c, code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		} else {
			c.JSON(http.StatusOK, token)
		}
	})

	r.Run(":8080")
}
