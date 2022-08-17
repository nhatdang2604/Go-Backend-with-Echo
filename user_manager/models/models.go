package models

type User struct {
	Id    int32  `orm:"auto" json:"id" form:"id"`
	Name  string `orm:"size(30)" json:"name" form:"name"`
	Age   int32  `json:"age" form:"age"`
	Phone string `orm:"size(11)" json:"phone" form:"phone"`
}

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
