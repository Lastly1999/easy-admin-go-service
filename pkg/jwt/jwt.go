package jwt

import (
	"easy-admin-go-service/helps"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	ID     string `json:"userId"`
	Name   string `json:"name"`
	RoleId string `json:"roleId"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//该路由下不校验token
	noVerify = []string{"/login", "/ping"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

// GenerateToken 生成token
func GenerateToken(claims *UserClaims) string {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

// JwtVerify 验证token
func JwtVerify(c *gin.Context) {
	//过滤是否验证token
	if helps.IsContainInArray(c.Request.RequestURI, noVerify) {
		return
	}
	token := c.GetHeader("Authorization")
	if token == "" {
		panic("token not exist !")
	}
	// 验证token，并存储在请求中
	err, u := ParseToken(token)
	if err != nil {
		return
	}
	c.Set("user", u)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (error, *UserClaims) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return err, nil
	}
	return nil, claims
}

// Refresh 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}
