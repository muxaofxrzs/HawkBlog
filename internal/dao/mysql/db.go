package mysql

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var GlobalConn sqlx.SqlConn

func NewDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=True", "root", "hawk123", "1.94.27.198", 3306, "hawk")
	conn := sqlx.NewMysql(dsn)
	GlobalConn = conn
}
