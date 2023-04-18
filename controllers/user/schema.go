package userctrl

// Schema
type CreateUserSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type LoginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
