package config

import (
    "fmt"
    "os"

    "github.com/go-ini/ini"
)

var (
	AppMode  string
	HttpPort string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// `go`中包的初始化顺序：`初始化包内声明的变量`、`init()`、`main()`
// init() 仅应用来初始化包内变量
func init() {
	cfg, err := ini.Load("./config/config.ini")
    if err != nil {
        fmt.Printf("Read config.ini err : [%v]", err)
        os.Exit(1)
    }

	convertConfig(cfg)
}

// @description  转化ini配置文件的数据
// @param config config.ini文件中的内容
func convertConfig(cfg *ini.File) {
	AppMode = cfg.Section("").Key("app_mode").String()
	HttpPort = cfg.Section("server").Key("http_port").String()
	DbHost = cfg.Section("database").Key("db_host").MustString("localhost")
	DbPort = cfg.Section("database").Key("db_port").MustString("3306")
	DbUser = cfg.Section("database").Key("db_user").MustString("ginblog")
	DbPassWord = cfg.Section("database").Key("db_password").String()
	DbName = cfg.Section("database").Key("db_name").MustString("ginblog")
}