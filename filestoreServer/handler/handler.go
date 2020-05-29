package handler

/*
@Time : 2020/5/27 9:38 下午
@Author : audiRS7
@File : handler.go
@Software: GoLand
*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "inter server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err: %s\n", err.Error())
			return
		}
		defer file.Close()
		newfile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Create file failed,err %s\n", err.Error())
			return
		}
		defer newfile.Close()

		_, err = io.Copy(newfile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file,err: %s\n", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}
