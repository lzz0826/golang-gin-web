package main

import (
	"github.com/gin-gonic/gin"
	"web/drivers"
	server "web/servers"
)

var HttpServer *gin.Engine

func main() {
	// 服务停止时清理数据库链接
	defer drivers.MysqlDb.Close()
	// 启动服务
	server.Run(HttpServer)

}
