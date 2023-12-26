package payload

type User struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserInfo struct {
	User User `json:"user"`
}
