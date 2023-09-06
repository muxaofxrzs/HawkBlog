package mysql

import (
	"context"
	"fmt"
	"hawk/model"
)

func GetUser(name, pwd string) (*model.User, int) {
	user := &model.User{}
	sql := "SELECT `Id`, `Name`, `UserName` FROM `user` WHERE `UserName` = ? AND `password` = ? LIMIT 1"
	err := GlobalConn.QueryRowCtx(context.Background(), user, sql, name, pwd)
	if err != nil {
		fmt.Printf("err:%s\n", err)

		return nil, 1
	}
	return user, 0
}
