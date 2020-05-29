package meta

/*
@Time : 2020/5/27 11:21 下午
@Author : audiRS7
@File : filemeta.go
@Software: GoLand
*/

//文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}
