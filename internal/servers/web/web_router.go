package web

import (
	"bito_group/internal/common/logs"
	"bito_group/internal/user"
)

func WithRouter(s *WebServer) {

	logs.Debugf("註冊路由")

	// 新建 handler
	userHandler := user.NewUserHandler(s.Apps.UserApp)

	// api
	api := s.Engin.Group("/v1")

	// 中间件
	//api.Use(authMiddleware.Auth)

	// 路由
	//api.GET("/UserInfo", userHandler.UserInfo)
	api.POST("/AddSinglePersonAndMatch", userHandler.AddSinglePersonAndMatch)
	api.DELETE("/RemoveSinglePerson", userHandler.RemoveSinglePerson)
	api.POST("/QuerySinglePeople", userHandler.QuerySinglePeople)
}
