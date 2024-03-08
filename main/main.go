package main

import (
	"ElectronicsRecycle/internal/handler"
	"log"
	"net/http"
)

func main() {
	// 设置路由处理函数来处理图片分类的API请求
	http.HandleFunc("/api/image/classify", handler.ClassifyImageHandler)

	// 设置静态文件服务，假设静态文件放在项目根目录的static文件夹
	fs := http.FileServer(http.Dir("static"))
	// 为根URL（"/"）设置文件服务器，使得访问根URL时能够看到index.html页面
	http.Handle("/", fs)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
