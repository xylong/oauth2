package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
	"oauth2/utils"
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

	r.POST("token", func(c *gin.Context) {
		if err := s.HandleTokenRequest(c.Writer, c.Request); err != nil {
			panic(err.Error())
		}
	})

	r.Any("login", func(c *gin.Context) {
		data := map[string]string{
			"error": "",
		}

		if c.Request.Method == http.MethodPost {
			name, pass := c.PostForm("userName"), c.PostForm("userPass")
			if name+pass == "张三123456" {
				utils.SaveUserSession(c, name)
				c.Redirect(http.StatusFound, "/auth?"+c.Request.URL.RawQuery)
				return
			} else {
				data["error"] = "用户名或密码错误"
			}
		}

		c.HTML(http.StatusOK, "login.html", data)
	})
	r.Run(":80")
}

// userAuthorizeHandler 用户授权
func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if userID = utils.GetUserSession(r); userID == "" {
		w.Header().Set("Location", "login?"+r.URL.RawQuery)
		w.WriteHeader(http.StatusFound)
	}
	return
}
