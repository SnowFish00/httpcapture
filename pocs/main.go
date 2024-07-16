package main

import (
	"demo/db"
	"demo/project/save"
	"demo/project/search"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	db.GetDB()
}

func main() {
	// 创建 Gin 引擎
	r := gin.Default()

	// 定义 /save 路由
	r.POST("/save", save.SavePoc)
	r.GET("/getpoc", search.GetPocByID)

	// 运行 HTTP 服务
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("无法启动服务器: %v", err)
	}
}
