package main

import (
	"bito_group/config"
	"bito_group/internal/common/logs"
	"bito_group/internal/common/signals"
	"bito_group/internal/servers"
	"bito_group/internal/servers/web"
)

// NewServers 通过配置文件初始化 Repo 依赖, 然后初始化 App, 最后组装为 Server
// 比如 UserRepo -> UserApp -> WebServer
func NewServers(cfg *config.SugaredConfig) servers.ServerInterface {

	repos := servers.NewRepos(cfg) // config 設定好 db
	apps := servers.NewApps(repos) // db 創建 usercase, 返回usecase

	servers := servers.NewServers()
	servers.AddServer(web.NewWebServer(cfg, apps)) // 啟動 web server

	return servers
}

func main() {

	// 初始化 config 配置
	cfg := config.NewConfig("./config.yaml")
	// 初始化日志
	logs.Init(cfg.Log)

	logs.Debugf("mode=%+v", cfg.Web)

	// 獲得 servers, 比如 WebServer, Websocket, RpcServer
	servers := NewServers(cfg)

	// 啟動 servers
	servers.AsyncStart()

	logs.Debugf("優雅關閉 等待訊號中...")
	signals.WaitWith(servers.Stop)

}
