package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var (
	sessionStore *sessions.CookieStore
	key="LoginUser"
	uid="userID"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if err:=viper.ReadInConfig();err!=nil {
		log.Fatal("read config failed: %v", err)
	}
	sessionStore=sessions.NewCookieStore([]byte(viper.Get("session.secret").(string)))
	sessionStore.Options.Domain="oauth.think.com"
	sessionStore.Options.Path= "/"
	sessionStore.Options.MaxAge=0	// 关掉浏览器就清session
}

// SaveUserSession 保存用户session
func SaveUserSession(ctx *gin.Context,userID string)  {
	s,err:=sessionStore.Get(ctx.Request, key)
	if err!=nil {
		panic(err.Error())
	}
	s.Values[uid]=userID
	err=s.Save(ctx.Request,ctx.Writer)
	if err!=nil {
		panic(err.Error())
	}
}

// GetUserSession 保存用户session
func GetUserSession(r *http.Request) string {
	if s,err:=sessionStore.Get(r,key);err!=nil {
		if s!=nil && s.Values[uid]!=nil {
			return s.Values[uid].(string)
		}
	}
	return ""
}