package mysql

import (
	"hawk/internal/types"
)

func UpdateUserInfo(req *types.UpdateUserInformationReq, userId int64) int {
	sql := "UPDATE user SET UserName = ?,PassWord = ?,Email = ?,Gender =?,Age = ?, Interest = ?, PhoneNumber =? WHERE Id = ?"
	_, err := GlobalConn.Exec(sql, req.UserName, req.PassWord, req.Email, req.Gender, req.Age, req.Interest, req.PhoneNumber, userId)
	if err != nil {
		return 1
	}
	return 0
}
