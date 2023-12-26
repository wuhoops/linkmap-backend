package database

type User struct {
	UserId   string `json:"user_id" gorm:"primaryKey; not null; unique"`
	Email    string `json:"email" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
	UserName string `json:"username" gorm:"unique"`
	Cards    []Card `json:"cards" gorm:"foreignKey:OwnerId; references:UserId"`
}
