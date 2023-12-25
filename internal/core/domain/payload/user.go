package payload

type User struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

type UserInfo struct {
	User User `json:"user"`
}
