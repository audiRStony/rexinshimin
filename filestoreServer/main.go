package main

import (
	"filestoreServer/handler"
	"fmt"
	"net/http"
)

/*
@Time : 2020/5/27 9:38 下午
@Author : audiRS7
@File : main.go
@Software: GoLand
*/

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server,err : %s\n", err.Error())
	}
}
