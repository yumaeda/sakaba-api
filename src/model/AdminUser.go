package model

type AdminUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
