package models

type Message struct {
	Text string
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" xml:"username" query: "username"`
	Password string `json:"password" form:"password" xml:"password" query: "password"`
}
