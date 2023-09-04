package middleware

import (
	"fmt"
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
		// Passthrough to next handler if need
		next(w, r)
	}
}
