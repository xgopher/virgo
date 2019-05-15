### 简介

 - 基于gorm的分页插件

### 分页传参

```
type Param struct {
	DB      *gorm.DB
	Page    int
	PerPage int
	OrderBy []string
}

```
Page: 当前页码
PerPage: 分页大小
OrderBy: 排序规则，同gorm写法

### 返回示例
```
{
    "meta": {
        "pagination": {
            "per_page": 20, // 分页大小
            "total": 5, // 当前页总条数
            "from": 1, // 当前页第一条记录为第几条，例如分页20条，当前页为2，则from为21
            "to": 1, // 当前页最后一条记录为第几条，例如分页20条，当前页为2，则from为40
            "current_page": 1, // 当前页码
            "last_page": 1, // 最后一页的页码
            "prev_page": 1,
            "next_page": 1
        }
    }
    "data": [
        {
            "id": 1,
            "firstname": "liqun",
            "lastname": "xu",
            "username": "test",
            "password": "123456",
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null
        }
    ],
}
```

### Gin中配合gorm使用Paginator

```golang
r := gin.Default()
r.GET("/", func(c *gin.Context) {
    var db *gorm.DB
    db = database.DB // 自己实现db连接
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    PerPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "0"))
    var users []User

    paginator := pagination.Pagging(&pagination.Param{
        DB:      db,
        Page:    page,
        PerPage: PerPage,
        OrderBy: []string{"id desc"},
        ShowSQL: false, // 是否启动debug
    }, &users)
    c.JSON(200, paginator)
})
```
