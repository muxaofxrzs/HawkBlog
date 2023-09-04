package mysql

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	//这里少了一句
)

var GlobalConn sqlx.SqlConn

func New() {
	//parseTime=True&loc=Local MySQL 默认时间是格林尼治时间，与我们差八小时，需要定位到我们当地时间
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "hawk123", "1.94.27.198:3306", "hawk")
	conn := sqlx.NewMysql(dsn)
	GlobalConn = conn
}
