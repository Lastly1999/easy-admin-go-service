package system_auth_service

type Auth struct {
	UserName string `json:"string"`
}

func (auth *Auth) CheckLogin() string {
	return "ok"
}
