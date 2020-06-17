package session

import (
	"demo_gin/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 初始化session
func InitSession(engine *gin.Engine) {
	cfg := config.GetConfig()
	store, err := redis.NewStore(10, "tcp", cfg.RedisHost+":"+cfg.RedisPort, "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	engine.Use(sessions.Sessions("session", store))
}

// set
func SetSession(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

// get
func GetSession(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
