package config

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Effect *casbin.Enforcer

func Rbac() {
	a, err := xormadapter.NewAdapter("mysql", "root:hawk123@tcp(1.94.27.198:3306)/hawk", true)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	//		m, err := model.NewModelFromString(`
	//[request_definition]
	//r = sub, obj, act
	//
	//[policy_definition]
	//p = sub, obj, act
	//
	//[policy_effect]
	//e = some(where (p.eft == allow))
	//
	//[matchers]
	//m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	//`)
	//		if err != nil {
	//			log.Fatalf("error: model: %s", err)
	//		}

	e, err := casbin.NewEnforcer("./internal/config/rbac_model.conf", a)
	Effect = e
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
}
