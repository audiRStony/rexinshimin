package utils

/*
@Time : 2020-03-07 18:14
@Author : audiRStony
@File : misc.go
@Software: GoLand
*/

/*
将url加上http://$ip:$port 前缀
http:// + 127.0.0.1 + : + 8080 + $request
*/

func AddDomain2Url(url string) (domain_url string)  {
    domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url //url为图片存放的路径

    return domain_url
}
