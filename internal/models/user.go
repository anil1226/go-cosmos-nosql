package models

type User struct {
	UserResp
	Password string `json:"password"`
}

type UserResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
