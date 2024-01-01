[English](README.md) 

# gorm-zero

## go zero gorm 拓展

### 如果你使用gozero框架,又想使用gorm访问数据库,你可以使用这个库

## 特性

- 可集成进gozero
- 可以使用goctl生成代码
- 默认使用gozero logx日志库
- 支持链路追踪
- 支持指标监控
- gorm库增强



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

# 示例

### 创建一个studentManager.sql文件并写入建表语句

```sql
CREATE TABLE `student`(
    `id` INT AUTO_INCREMENT,
    `name` varchar(10) NOT NULL DEFAULT '',
    `age` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)charset=utf8mb4;

CREATE TABLE `course`(
    `id` INT AUTO_INCREMENT,
    `name` varchar(20) NOT NULL DEFAULT '',
    `credit` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)charset=utf8mb4;

CREATE TABLE `sc` (
    `id` INT AUTO_INCREMENT,
    `student_id` INT NOT NULL,
    `course_id` INT NOT NULL,
    `grade` INT NOT NULL,
    PRIMARY KEY (`id`)
);
```

### 生成代码

```shell
goctl template init --home ./template
cd template/model
go run github.com/klen-ygs/gorm-zero/model@latest
cd ../.. 
goctl model mysql ddl -dir ./internal/model -src studentManager.sql -style go_zero -home template
go mod tidy
```

### 在 student_model.go 中编写你的代码

```go
// customStudentLogicModel 在这里编写你的代码
customStudentLogicModel interface {
    WithSession(tx *gorm.DB) StudentModel

    // FindStudentByName 自己的db方法
    FindStudentByName(ctx context.Context, name string) ([]*Student, error)
    // FindStudentAgeBetween 自己的db方法
	FindStudentAgeBetween(ctx context.Context, start, ent int) ([]*Student, error)
    // FindStudentGradeGreater 自己的db方法
	FindStudentGradeGreater(ctx context.Context, greater int) (resp []*Student,err error)
}

// FindStudentByName select * from student where name = #{name}
func (c customStudentModel) FindStudentByName(ctx context.Context, name string) (resp []*Student, err error) {
	err = c.conn.
		WithContext(ctx).
		Where(Eq(&QStudent.Name, name)).
		Find(&resp).
		Error

	return
}
	
// FindStudentAgeBetween select * from student where age between #{start} and #{end}
func (c customStudentModel) FindStudentAgeBetween(ctx context.Context, start, ent int) (resp []*Student,err error) {
	err = c.conn.
		WithContext(ctx).
		Where(Between(&QStudent.Age, start, ent)).
		Find(&resp).
		Error
	
	return 
}

// FindStudentGradeGreater
// select student.* from student as stu
// join sc on stu.id = sc.student_id
// where sc.grade > #{greater}
func (c customStudentModel) FindStudentGradeGreater(ctx context.Context, greater int) (resp []*Student, err error) {
	err = c.conn.
		WithContext(ctx).
		Joins(On(&QSc, &QSc.StudentId, &QStudent.Id)).
		Where(Gt(&QSc.Grade, greater)).
		Find(&resp).
		Error

	return
}

```

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


