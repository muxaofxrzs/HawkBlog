package main

import (
	"flag"
	"fmt"
	"hawk/global"
	"hawk/internal/config"
	"hawk/internal/dao/mongo"
	"hawk/internal/dao/mysql"
	"hawk/internal/dao/redis"
	"hawk/internal/handler"
	"hawk/internal/handler/im"
	"hawk/internal/pkg/snowflake"
	"hawk/internal/svc"
	"hawk/internal/tools"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hawk.yaml", "the config file")

func init() {
	err := global.SetupCasbinEnforcer()
	if err != nil {
		fmt.Print("rbac预加载失败")
		//log.Fatalf("init.SetupCasbinEnforcer err: %v", err)
		//global.Logger.Fatalf("init.SetupCasbinEnforcer err: %v", err)
	}
}
func main() {
	flag.Parse()
	config.Rbac()
	mysql.New()
	redis.CreateRedisClient()
	mongo.NewMongo()
	defer func() {
		mongo.CloseMongo()
	}()

	snowflake.Init("2023-08-30", 1)
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	go func() {
		//chatRoom := im.NewChatRoom()
		http.HandleFunc("/ws", tools.Token.ImJwtAuthMiddleware(im.HandIeWebSocket))
		log.Println("IM Websocket Starting server at: 9090...")
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("聊天服务启动失败！！！:", err)
		}
	}()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
