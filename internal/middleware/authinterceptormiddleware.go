package middleware

import (
	"fmt"
	"hawk/internal/config"
	"hawk/internal/tools"
	"net/http"
)

type AuthInterceptorMiddleware struct {
}

func NewAuthInterceptorMiddleware() *AuthInterceptorMiddleware {
	return &AuthInterceptorMiddleware{}
}

func (m *AuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("111111111111111111")
		// TODO generate middleware implement function, delete after code implementation
		debugHeaderValue := r.Header.Get("debug")

		if debugHeaderValue != "" {
			// 如果 "debug" 头部不为空，跳过后续中间件和处理器
			return
		}
		fmt.Println("22222222222222222222")

		//debug是什么？
		auth := r.Header.Get("Authorization")
		//auth = auth[7:]
		fmt.Println("33333333333333333")

		fmt.Printf("auth:%+v\n", auth)

		data, err := tools.Token.VerifyToken(auth)
		if err != nil {
			fmt.Println("jwt解析出问题l")
		}
		fmt.Printf("data:%+v\n", data)

		obj := r.URL.Path
		// 请求的方法
		act := r.Method

		sub := data.Role

		//sub := "alice" // 想要访问资源的用户。
		//obj := "data1" // 将被访问的资源。
		//act := "read"  // 用户对资源执行的操作。

		ok, err := config.Effect.Enforce(sub, obj, act)

		if err != nil {
			// 处理err
		}

		if ok == true {
			fmt.Println("成功")
		} else {
			return
			// 拒绝请求，抛出异常
		}

		// 您可以使用BatchEnforce()来批量执行一些请求
		// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
		// 例如results[0] 是{"alice", "data1", "read"}的结果
		_, err = config.Effect.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})

		next(w, r)
	}
}
