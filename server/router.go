package server

import (
	"gomusic/api"
	"gomusic/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	//固定路由
	apiHttpPrefix := "/api/v1"

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	rg := r.Group(apiHttpPrefix)

	registerDebug(rg)
	registerLogin(rg)
	registerExample(rg)
	registerMusic(rg)

	return r
}

func registerDebug(rg *gin.RouterGroup) {
	// Ping
	rg.POST("/debug/ping", api.Ping)

	// Version
	rg.GET("/debug/version", api.Version)
}

func registerLogin(rg *gin.RouterGroup) {

	// 用户注册
	rg.POST("user/register", api.UserRegister)

	// 用户登录
	rg.POST("user/login", api.UserLogin)

	// 需要登录保护的
	auth := rg.Group("")
	auth.Use(middleware.AuthRequired())
	{
		// User Routing
		auth.GET("user/me", api.UserMe)
		auth.DELETE("user/logout", api.UserLogout)
	}
}

func registerExample(rg *gin.RouterGroup) {
	egr := rg.Group("/example")

	egr.GET("/list", api.GetExampleList) // 获取样例列表
	egr.GET("", api.GetExample)          // 获取样例
	egr.POST("", api.PostExample)        // 添加样例
	egr.PUT("", api.PutExample)          // 修改样例
	egr.DELETE("", api.DeleteExample)    // 删除样例
}

func registerMusic(rg *gin.RouterGroup) {
	egr := rg.Group("/music")

	egr.GET("", api.GetMusic)              // 获取单首歌曲列表
	egr.GET("/list", api.GetMusicList)     // 获取歌曲列表
	egr.GET("/random", api.GetRandomMusic) // 获取随机若干首歌曲

	// 需要登录保护的
	auth := egr.Group("")
	auth.Use(middleware.AuthRequired())
	{
		auth.POST("", api.PostMusic)              // 添加歌曲
		auth.PUT("", api.PutMusic)                // 修改歌曲
		auth.DELETE("", api.DeleteMusic)          // 删除歌曲
		auth.POST("/list", api.PostBatchMusic)    // 批量添加歌曲
		auth.DELETE("list", api.DeleteBatchMusic) // 批量删除歌曲
	}
}
