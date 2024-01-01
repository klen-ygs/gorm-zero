[中文](README-zh.md)

# gorm-zero

## go zero gorm extension

### If you use go zero, and you want to use Gorm. You can use this library.

## Feature

- Integrate go-zero
- generate code by goctl
- use logx default
- support opentelemetry
- support  metrics
- more easily to use gorm



# Usage

### add the dependent

```shell
go get github.com/klen-ygs/gorm-zero
```
### generate code

you can generate code by three options.

1. auto replace

   ```shell
   goctl template init --home ./template
   cd template/model
   go run github.com/klen-ygs/gorm-zero/model@latest
   ```

2. download model template to local

- replace  template/model in your project with gorm-zero/model

- generate

  ```shell
  goctl model mysql -src={patterns} -dir={dir} -cache --home ./template
  ```

3. generate by remote template

set remote = https://github.com/klen-ygs/gorm-zero.git

```shell
goctl model mysql -src={patterns} -dir={dir} -cache --remote https://github.com/klen-ygs/gorm-zero.git
```



## Mysql

### Config
```go
import (
    "github.com/klen-ygs/gorm-zero/gormc/config/mysql"
)
type Config struct {
    Mysql mysql.Conf
    // ...
}
```
### Initialization

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

or

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

### Config
```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/pg"
)
type Config struct {
    PgSql pg.Conf
    // ...
}
```
### Initialization

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

or

```go
import (
"github.com/klen-ygs/gorm-zero/gormc/config/pg"
)
func NewServiceContext(c config.Config) *ServiceContext {
    db := pg.MustConnect(c.PgSql)
    // ...
}
```

# Example

### create sql file named studentManager.sql 

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

### generate code

```shell
goctl template init --home ./template
cd template/model
go run github.com/klen-ygs/gorm-zero/model@latest
cd ../.. 
goctl model mysql ddl -dir ./internal/model -src studentManager.sql -style go_zero -home template
go mod tidy
```

### write code in student_model.go

```go

// customStudentLogicModel add your method here
customStudentLogicModel interface {
    WithSession(tx *gorm.DB) StudentModel

    // FindStudentByName my db method
    FindStudentByName(ctx context.Context, name string) ([]*Student, error)
    // FindStudentAgeBetween my db method
	FindStudentAgeBetween(ctx context.Context, start, ent int) ([]*Student, error)
    // FindStudentGradeGreater my db method
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

### Transition

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

