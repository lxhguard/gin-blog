# G004 go-ini使用

# 一.安装

通常配置文件不会在代码里硬编码，通常是放在配置文件中的。

常见的配置文件有很多，比如 json 、 xml 、 ini 等等。

go-ini 是 Go 语言中用于操作 ini 文件的第三方库。

```js
// 可以使用最新版本的 go-ini （推荐使用这个）
$  go get github.com/go-ini/ini
// 如果想拉取指定版本的 go-ini
$ go get gopkg.in/ini.v1
```

# 二.使用

[go-ini官方文档](https://ini.unknwon.cn/)

ini 有四个概念：分区（Section），键（key），值（value），注释（Comment）。

```ini
# possible values : production, development
app_mode = development

# 这是第一个分区，默认分区
# 要想获取 app_mode 的值 ，可通过 cfg.Section("").Key("app_mode").String()
# .String() 含义是获取一个 string 类型的值

[server]
# Protocol (http or https)
protocol = http

# 这是第二个分区，名为 server
# 要想获取 protocol 的值 ，可通过 cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})
# In() 方法代表了，取出来的字符串一定是 "http"、 "https" 中的某一个。
```

看到上面的 cfg 你可能会有困惑，这个变量从哪来的， 它其实就是使用了 go-ini 的 load 方法，具体如下：

```go
import (
    "github.com/go-ini/ini"
)
func main() {
    cfg, err := ini.Load("my.ini")
}
```

`.Key()` 之后可以链式调用很多方法，除了 `.String()` 、`.In()` ，还有 `.Value()` 可以直接获取原值（这种方式性能最佳）。

比较经典的还有 `.MustXxx` 系列。比如：

由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，当键不存在或者转换失败时，则会直接返回该默认值。但是，MustString 方法必须传递一个默认值。

# 三.实战

这里我们只需要会 配置ini文件，会使用ini的读取API 即可，就可以直接用来玩项目了。

至于别的API，感兴趣的可以自己看，但通常项目用不上。

`go`中包的初始化顺序：`初始化包内声明的变量`、`init()`、`main()`。

看下我们项目仓库的 `go.mod` 文件，先知道当前项目的`module ginblog`，名字是`ginblog`，所以当前项目下的包依赖应该使用`ginblog`。

```go
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
```




