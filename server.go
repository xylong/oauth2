package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
)

func main() {
	// 1.创建管理对象
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	// 2.配置网站信息
	clientStore := store.NewClientStore()
	err := clientStore.Set("clienta", &models.Client{
		ID:     "clienta",
		Secret: "123",
		Domain: "http://127.0.0.1:8080",
	})
	if err != nil {
		log.Fatal(err)
	}
	manager.MapClientStorage(clientStore)
	// 3.
	s := server.NewDefaultServer(manager)
	s.SetUserAuthorizationHandler(userAuthorizeHandler)
	r := gin.New()
	r.LoadHTMLGlob("public/*.html")
	// 响应授权码
	r.GET("auth", func(c *gin.Context) {
		err := s.HandleAuthorizeRequest(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
		}
	})
	r.GET("login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.Run(":80")
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	w.Header().Set("Location", "login")
	w.WriteHeader(http.StatusFound)
	return
}
