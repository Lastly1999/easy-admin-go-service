package system_auth_service

import (
	"easy-admin-go-service/app/services/admin/response"
	"easy-admin-go-service/app/services/admin/system_role_service"
	"easy-admin-go-service/app/services/admin/system_user_service"
	"easy-admin-go-service/pkg/captcha"
	"easy-admin-go-service/pkg/jwt"
	"errors"
	uuid "github.com/satori/go.uuid"
	"log"
	"strconv"
)

type AuthService struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	SmsUuid  string `json:"smsUuid"`
	SmsCode  string `json:"smsCode"`
}

type IAuthService interface {
	CheckLogin() (error, *response.CheckLoginResponse)
	GenerateGraphicCode() (error, *response.GenerateGraphicCodeResponse)
	FindUserRoleMenus() (error, *response.GetRoleMenusResponse)
}

// CheckLogin 验证登录
func (authService *AuthService) CheckLogin() (error, *response.CheckLoginResponse) {
	userService := &system_user_service.UserService{
		UserName: authService.UserName,
		PassWord: authService.PassWord,
	}
	resp, err := userService.FindUserInfo()
	// 查询是否存在用户
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
func (authService *AuthService) GenerateGraphicCode() (error, *response.GenerateGraphicCodeResponse) {
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

// FindUserRoleMenus 获取用户权限菜单信息
func (authService *AuthService) FindUserRoleMenus() (error, *response.GetRoleMenusResponse) {
	service := system_role_service.RoleService{}
	err, roleIds := service.GetUserRoleIds(1)
	log.Panicln(roleIds)
	if err != nil {
		return err, nil
	}
	return nil, nil
}
