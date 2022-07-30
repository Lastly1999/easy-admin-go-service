package admin

import (
	"easy-admin-go-service/app/services/admin/request"
	"easy-admin-go-service/app/services/admin/system_user_service"
	"easy-admin-go-service/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserApi struct {
}

// List 获取系统用户列表
// @Summary 获取系统用户列表
// @Description 获取系统用户列表
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Router /user/user [posts]
func (userApi *UserApi) List(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	userSrv := system_user_service.UserService{}
	list, err := userSrv.FindUserList()
	if err != nil {
		rep.Error(http.StatusInternalServerError, "参数转换异常")
		return
	}
	rep.Success(list)
}

// Create 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Router /user/user [put]
func (userApi *UserApi) Create(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	createUserRequest := &request.CreateUserRequest{}
	err := rep.Ctx.ShouldBindJSON(createUserRequest)
	if err != nil {
		rep.Error(http.StatusInternalServerError, "请求参数解析异常")
		return
	}
	useSrv := system_user_service.UserService{}
	err = useSrv.InsertUser(createUserRequest)
	if err != nil {
		rep.Error(http.StatusInternalServerError, err.Error())
		return
	}
	rep.Success(nil)
}

// Update 更新用户
// @Summary 更新用户
// @Description 更新用户
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Router /user/user [patch]
func (userApi *UserApi) Update(ctx *gin.Context) {
	rep := global.NewResult(ctx)
	updateUserRequest := &request.UpdateUserRequest{}
	err := rep.Ctx.ShouldBindJSON(updateUserRequest)
	if err != nil {
		rep.Error(http.StatusInternalServerError, "参数解析失败")
		return
	}
	userSrv := system_user_service.UserService{}
	err = userSrv.UpdateUser(updateUserRequest)
	if err != nil {
		rep.Error(http.StatusInternalServerError, err.Error())
		return
	}
}
