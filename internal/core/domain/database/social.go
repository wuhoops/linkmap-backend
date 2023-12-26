package database

type Topic string

const (
	Instagram Topic = "Instagram"
	Facebook  Topic = "Facebook"
	Line      Topic = "Line"
	Linkedin  Topic = "Linkedin"
	Twitter   Topic = "Twitter"
	Whatsapp  Topic = "Whatsapp"
)

type Social struct {
	SocialId string `json:"social_id" gorm:"primaryKey"`
	OwnerId  string `json:"owner_id" gorm:"not null"`
	Topic    Topic  `json:"topic" gorm:"not null"`
	Link     string `json:"link" gorm:"not null"`
}
