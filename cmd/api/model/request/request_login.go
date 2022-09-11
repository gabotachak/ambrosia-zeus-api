package request

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
