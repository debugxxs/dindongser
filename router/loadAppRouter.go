package router

import (
	"dingdongser/controller"
	"dingdongser/middleware/cors"
	"dingdongser/middleware/jwt"
	"github.com/gin-gonic/gin"
)

var (
	myJwt jwt.AuthJwt
	user controller.UserController
)

func LoadRouter(engine *gin.Engine) {
	allAuthMiddleware := myJwt.AuthMiddlewareFunc(myJwt.AllAuthMiddleware)
	engine.Use(cors.Cors())
	engine.NoRoute(allAuthMiddleware.MiddlewareFunc(), myJwt.NoRouteHandler)
	engine.POST("/login", allAuthMiddleware.LoginHandler)
	userApi := engine.Group("/user")
	{
		userApi.GET("/refresh_token", allAuthMiddleware.RefreshHandler)
	}
	userApi.Use(allAuthMiddleware.MiddlewareFunc())
	{
		userApi.POST("problem",user.SetNewProblem)
		userApi.PUT("/password",user.ModifyPass)
	}
}
