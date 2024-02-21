package test

import (
	"github.com/klen-ygs/gorm-zero/gormc/config/mysql"
	"github.com/zeromicro/go-zero/core/conf"
	"strings"
	"testing"
)

type Conf struct {
	Mysql mysql.Conf
}

func TestConnMysql(t *testing.T) {
	var c Conf

	conf.MustLoad("./myconf.yaml", &c)

	_, err := mysql.Connect(c.Mysql)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTableName(t *testing.T) {
	trim := strings.Trim("`table`", "`")
	if strings.HasSuffix(trim, "`") || strings.HasPrefix(trim, "`") {
		t.Fatal("删除``失败")
	}
}
