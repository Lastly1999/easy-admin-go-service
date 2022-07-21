package middlewares

import (
	"easy-admin-go-service/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		whiteRouter := []string{
			"/v1/auth/login",
			"/v1/auth/graphic",
			"/swagger/index.html",
		}
		for _, v := range whiteRouter {
			if v == context.Request.URL.Path {
				// 白名单之内 放行
				context.Next()
			}
		}
		if len(auth) == 0 {
			context.Abort()
			context.String(http.StatusOK, "未登录⽆权限")
			return
		}
		// 只要出错直接拒绝请求 校验 token
		err, _ := jwt.ParseToken(auth)
		if err != nil {
			context.Abort()
			message := err.Error()
			context.JSON(http.StatusOK, message)
			return
		} else {
			println("token 正确")
		}
	}
}
