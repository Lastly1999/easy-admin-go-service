package request

type CheckLoginRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}
