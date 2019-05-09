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


## 参考

- http://www.jyguagua.com/?p=3038