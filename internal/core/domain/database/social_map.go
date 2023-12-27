package database

type SocialMap struct {
	SocialId string `json:"social_id" gorm:"not null; unique"`
	Name     string `json:"name" gorm:"not null"`
}
