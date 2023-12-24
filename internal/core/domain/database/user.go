package database

type User struct {
	UserId   string `json:"user_id" gorm:"primaryKey;not null"`
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Cards    []Card `json:"cards" gorm:"foreignKey:OwnerID;references:UserId"`
}
