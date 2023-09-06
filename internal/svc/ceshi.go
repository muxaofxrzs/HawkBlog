package svc

//
//import (
//	"context"
//	"github.com/your/module/path/model"
//	"github.com/your/module/path/service/svc"
//	"github.com/zeromicro/go-zero/core/stores/sqlx"
//)
//
//type ServiceContext struct {
//	c      context.Context
//	svcCtx *svc.ServiceContext
//	conn   sqlx.SqlConn
//}
//
//func NewServiceContext(c context.Context, svcCtx *svc.ServiceContext) *ServiceContext {
//	return &ServiceContext{
//		c:      c,
//		svcCtx: svcCtx,
//		conn:   svcCtx.Database.SqlConn,
//	}
//}
//
//func (c *ServiceContext) GetUsersWithName(name string) ([]*model.User, error) {
//	query := "SELECT * FROM users WHERE name = ?"
//	var users []*model.User
//	err := c.conn.QueryRows(c.c, query, &users, name)
//	if err != nil {
//		return nil, err
//	}
//	return users, nil
//}
