package payload

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
