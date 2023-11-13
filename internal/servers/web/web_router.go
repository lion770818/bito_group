package web

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/user"
)

func WithRouter(s *WebServer) {

	logs.Debugf("註冊路由")

	// 新建 handler
	userHandler := user.NewUserHandler(s.Apps.UserApp)
	//authMiddleware := user.NewAuthMiddleware(s.Apps.UserApp)

	// 驗證
	// auth := s.Engin.Group("/auth")
	// auth.POST("/login", userHandler.Login)
	// auth.POST("/register", userHandler.Register)

	// api
	api := s.Engin.Group("/v1")

	// 中间件
	//api.Use(authMiddleware.Auth)

	// 路由
	api.GET("/user_info", userHandler.UserInfo)
	api.POST("AddSinglePersonAndMatch", userHandler.Register)
	api.DELETE("RemoveSinglePerson", userHandler.RemoveSinglePerson)
}
