package utils
/*
@Time : 2020-03-07 17:16
@Author : audiRStony
@File : config
@Software: GoLand
*/
import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/config"//使用beego框架的配置文件读取模块
)

var   (
    G_server_name string //项目名称
    G_server_addr string//服务器地址
    G_server_port string//服务端口
    G_redis_addr string//redis地址
    G_redis_port string//redis 端口
    G_redis_dbnum string//redis db编码
    G_mysql_addr string//mysql地址
    G_mysql_port string//mysql端口
    G_mysql_dbname string//mysql数据库名称
    G_fastdfs_port string//fastdfs 端口
    G_fastdfs_addr string//fastdfs地址
)

func InitConfig()  {
    //read configuration from configFile
    appconf,err := config.NewConfig("ini","./conf/app.conf")
    if err !=nil{
        beego.Debug(err)
        return
    }

    G_server_name = appconf.String("appname")
    G_fastdfs_addr = appconf.String("httpaddr")
    G_server_port = appconf.String("httpport")
    G_redis_addr = appconf.String("redisaddr")
    G_redis_port = appconf.String("redisport")
    G_redis_dbnum = appconf.String("redisdbnum")
    G_mysql_addr = appconf.String("mysqladdr")
    G_mysql_port = appconf.String("mysqlport")
    G_mysql_dbname = appconf.String("mysqldbname")
    G_fastdfs_addr = appconf.String("fastdfsaddr")
    G_fastdfs_port = appconf.String("fastdfsport")
    return
}

func init()  {
    InitConfig()
}