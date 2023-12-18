[English](README.md) 

# gorm-zero

## go zero gorm 拓展

### 如果你使用gozero框架,又想使用gorm访问数据库,你可以使用这个库

## 特性

- 可集成进gozero
- 可以使用goctl生成代码
- 默认使用gozero logx日志库
- 支持链路追踪



# 使用

### 下载依赖

```shell
go get github.com/klen-ygs/gorm-zero
```

### 生成代码

你可以通过以下三种方法生成代码

1. 自动替换模板

```shell
goctl template init --home ./template
cd template/model
go run github.com/klen-ygs/gorm-zero/model@latest
```

2. 下载model模板文件替换本地的模板

- 下载model文件夹替换你项目中的 template/model 

- 生成代码

  ```shell
  goctl model mysql -src={patterns} -dir={dir} -cache --home ./template
  ```

3. 使用远程仓库的模板文件

设置参数 remote = https://github.com/klen-ygs/gorm-zero.git

```shell
goctl model mysql -src={patterns} -dir={dir} -cache --remote https://github.com/klen-ygs/gorm-zero.git
```



## Mysql

### 配置

```go
import (
    "github.com/klen-ygs/gorm-zero/gormc/config/mysql"
)
type Config struct {
    Mysql mysql.Conf
    // ...
}
```

### 初始化

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/mysql"
)
func NewServiceContext(c config.Config) *ServiceContext {
    db, err := mysql.Connect(c.Mysql)
    if err != nil {
        log.Fatal(err)
    }
    // ...
}
```

或

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/mysql"
)
func NewServiceContext(c config.Config) *ServiceContext {
    db := mysql.MustConnect(c.Mysql)
    // ...
}
```



## PgSql

### 配置

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/pg"
)
type Config struct {
    PgSql pg.Conf
    // ...
}
```

### 初始化

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/pg"
)
func NewServiceContext(c config.Config) *ServiceContext {
    db, err := pg.Connect(c.PgSql)
    if err != nil {
        log.Fatal(err)
    }
    // ...
}
```

或

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/pg"
)
func NewServiceContext(c config.Config) *ServiceContext {
    db := pg.MustConnect(c.PgSql)
    // ...
}
```

# Coding

### 事务

```go
// use gormc.Transition, DB is *grom.DB
err = gormc.Transition(l.ctx, l.svcCtx.DB, func(tx *gorm.DB) (err error) {

    // use .WithSession 
    err = l.svcCtx.DepartmentsModel.WithSession(tx).
        Update(l.ctx, &model.Departments{
            DepartmentsName: "xxx",
        })

    return
})
if err != nil {
    return nil, err
}
```



### Query With Cache And Custom Expire Duration

```go
    gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdExpirePrefix, id)
    var resp Users
    err := m.QueryWithExpireCtx(ctx, &resp, gormzeroUsersIdKey, expire, func(conn *gorm.DB, v interface{}) error {
        return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
    })
    switch err {
        case nil:
            return &resp, nil
        case gormc.ErrNotFound:
            return nil, ErrNotFound
        default:
            return nil, err
    }
```

### Query With Cache And Default Expire Duration

```go
    gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdPrefix, id)
    var resp Users
    err := m.QueryCtx(ctx, &resp, gormzeroUsersIdKey, func(conn *gorm.DB, v interface{}) error {
        return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
    })
    switch err {
        case nil:
            return &resp, nil
        case gormc.ErrNotFound:
            return nil, ErrNotFound
        default:
            return nil, err
    }
```

