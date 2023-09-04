package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hawk/internal/tools"

	"hawk/internal/types"
)

func Registered(req *types.RegisterReq) int {
	user := &types.LoginRequest{}
	sql := "SELECT * FROM user WHERE UserName = ?"
	err := GlobalConn.QueryRowCtx(context.Background(), user, sql, req.Email)
	if err == nil {
		return 1
	}
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Error(err)
		return 2

	}

	//var worker *Worker
	worker := tools.NewWorker(001, 002)
	//ID:=gg.NextID()
	newId, _ := worker.NextID() // 使用雪花算法生成新的Id
	fmt.Println("newId:")
	fmt.Println(newId)

	sql = "INSERT INTO user (`Id`, `UserName`, `PassWord`, `Name`, `Email`, `Gender`, `Age`,`Interest`,`PhoneNumber`) VALUES (?, ?, ?, ?, ?, ?, ?,?,?)"
	r, err := GlobalConn.ExecCtx(context.Background(), sql, newId, req.UserName, req.PassWord, req.Name, req.Email, req.Gender, req.Age, req.Interest, req.PhoneNumber)
	if err != nil {
		panic(err)
		return 2
	}
	fmt.Println(r.RowsAffected())
	return 0
}
