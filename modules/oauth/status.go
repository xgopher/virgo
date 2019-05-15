// Package oauth 状态码定义
package oauth

// 状态码定义: 组成 = 模块编号(3位) + 错误级别(2位) + 错误编号(3位)
// 示例: 111 22 333,
// 注意: 模块编号不要和其它模块重复
const (
	StatusContinue = 11122333
)

var statusText = map[int]string{
	StatusContinue: "Continue",
}

// StatusText ...
func StatusText(code int) string {
	return statusText[code]
}
