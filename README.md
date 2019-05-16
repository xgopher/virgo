# Virgo


## 安装

- 线上部署及docker环境下：
    ```
    export GOPROXY=https://goproxy.io
    go build ./...
    ./app
    ```
- 本地开发(非docker环境下)：
    ```
    export GOPROXY=https://goproxy.io
    go mod tidy
    go run main.go
    ```
- 本地开发热更新（docker中不生效）

    - 下载bee工具：`go get github.com/beego/bee` (建议不要在此文件夹下执行，不然会进入mod文件里面)
    - 在此文件夹下执行`bee run`，会自动监听文件变化，重新编译

- 本地开发热更新2 (可能比bee好一点, 自主选择)
    - 下载fresh工具: `go get github.com/pilu/fresh` (不要在此文件夹下执行，不然会进入mod文件里面)
    - 在项目根目录下， 执行 `fresh` (对应配置文件 `runner.conf`), 会自动监听文件变化，重新编译
    - 开源项目地址：https://github.com/gravityblast/fresh
    - **懒人方法(里面写好了自动安装fresh)**：
        - 直接执行 `./run-fresh.sh`    

- 生产环境热更新 (详细使用，后续补全)

    - https://github.com/howeyc/fsnotify

## 单元测试

```
godotenv -f .env go test ./...
```

## 技术选型

- [x] 1. [Gin Framework](https://github.com/gin-gonic/gin)
- [x] 2. [Gin middleware for session management](https://github.com/gin-contrib/sessions) 支持 redis, memcache...
- [x] 3. [weixin/wechat/微信公众平台/微信企业号/微信商户平台/微信支付 go/golang sdk](https://github.com/chanxuehong/wechat)
- [x] 4. [ORM 数据库 [jinzhu/gorm]](https://github.com/jinzhu/gorm)
- [x] 5. [facebook ioc 依赖注入包](https://github.com/facebookarchive/inject)
- [x] 6. [gomodule/redigo 包](https://github.com/gomodule/redigo)

详情请看 `go.mod` 文件

## AJAX OPTIONS

如果您使用Javacript和CORS使用 `XMLHttpRequest` 或 `Fetch` ，则需要使用 `OPTIONS` 作为POST，PUT，DELETE请求。

首先，您必须添加2条路由。

```
v1.OPTIONS("/users", OptionsUser)      // POST
v1.OPTIONS("/users/:id", OptionsUser)  // PUT, DELETE
```

并声明 `OptionsUser` 函数：

```
func OptionsUser(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    c.Next()
}
```

## 请求参数

> @https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/

```
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
```

## 参考资料

- [Gin Framework](https://github.com/gin-gonic/gin)
- [Gin 官方文档](https://gin-gonic.com/zh-cn/docs/)
- [基于Go框架Gin开发RESTFul API](http://www.jyguagua.com/?p=3038)