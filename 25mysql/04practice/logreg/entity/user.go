package entity

type User struct {
	Id       int `json:"id,omitempty"`
	Age      int `json:"age,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
