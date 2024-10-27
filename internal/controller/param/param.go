package param

type Resp struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
type TokenResponse struct {
	Token   string `json:"token"`
	Expires int    `json:"expires_in"`
}
