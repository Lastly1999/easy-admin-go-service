package response

type GenerateGraphicCodeResponse struct {
	Uuid    string `json:"uuid"`
	BaseImg string `json:"baseImg"`
}

type CheckLoginResponse struct {
	Token string `json:"token"`
}
