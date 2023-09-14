package mysql

import (
	"context"
	"fmt"
	"hawk/model"
)

func GetUserInfo(userId int64) *model.UserInfo {
	userInfo := &model.UserInfo{}
	sql := "SELECT `Id`, `Name`, `UserName`, `PassWord`,`Name`,`Email`,`Gender`,`Age`,`Interest`, `PhoneNumber` FROM `user` WHERE `Id` = ? LIMIT 1"
	err := GlobalConn.QueryRowCtx(context.Background(), userInfo, sql, userId)
	if err != nil {
		fmt.Printf("err:%s\n", err)

		return nil
	}
	return userInfo
}
