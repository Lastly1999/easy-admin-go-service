package system_auth_service

import (
	"easy-admin-go-service/app/services/admin/response"
	"easy-admin-go-service/app/services/admin/system_user_service"
	"easy-admin-go-service/pkg/captcha"
	"easy-admin-go-service/pkg/jwt"
	"errors"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type Auth struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	SmsUuid  string `json:"smsUuid"`
	SmsCode  string `json:"smsCode"`
}

// CheckLogin 验证登录
func (auth *Auth) CheckLogin() (error, *response.CheckLoginResponse) {
	userService := &system_user_service.User{
		UserName: auth.UserName,
		PassWord: auth.PassWord,
	}
	// 查询是否存在用户
	err, resp := userService.GetUser()
	if err != nil {
		return err, nil
	}
	// 不存在用户
	if resp.CommonModel.ID == 0 {
		return errors.New("用户不存在"), nil
	}
	// 生成jwt的验证存储
	userClaims := jwt.UserClaims{
		ID:     strconv.Itoa(int(resp.ID)),
		Name:   resp.UserName,
		RoleId: strconv.Itoa(int(resp.RoleId)),
	}
	// 生成token
	token := jwt.GenerateToken(&userClaims)
	// 创建响应
	checkLoginResponse := &response.CheckLoginResponse{
		Token: token,
	}
	return nil, checkLoginResponse
}

// GenerateGraphicCode 生成图形验证码 存储redis
func (auth *Auth) GenerateGraphicCode() (error, *response.GenerateGraphicCodeResponse) {
	// 生成图形验证码
	answer, item := captcha.GenerateCaptchaHandler()
	// 生成uuid
	ids := uuid.NewV4().String()
	// 存储redis
	err := captcha.Base64CaptchaStore.Set(ids, answer)
	if err != nil {
		return err, nil
	}
	// 构造响应
	generateGraphicCodeResponse := &response.GenerateGraphicCodeResponse{
		Uuid:    ids,
		BaseImg: item.EncodeB64string(),
	}
	return nil, generateGraphicCodeResponse
}
