### 简介

 - Golang Gorm中自定义Time类型的JSON字段格式： [参考链接](http://www.axiaoxin.com/article/241/)

### 用法

- 原代码不需要更改太多，模型中`import "time"` 的地方改为 `import "app/services/time"`
```
import "app/services/time"

type Model struct {
	ID        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
```

### 说明
- 这个包尽量用在模型里面CreatedAt等字段，其他地方不建议使用

- 由于test是不走main.go的原因，本来时间格式应该从配置文件中获取，但是为了测试简单，这里就硬编码了，一般需要修改时间格式的情况也不多