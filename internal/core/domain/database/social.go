package database

import (
	"encoding/json"
	"fmt"
)

type Topic string

const (
	Instagram Topic = "instagram"
	Facebook  Topic = "facebook"
	Line      Topic = "line"
	Linkedin  Topic = "linkedin"
	X         Topic = "x"
	Whatsapp  Topic = "whatsapp"
)

type Social struct {
	SocialId string `json:"social_id" gorm:"not null; unique"`
	OwnerId  string `json:"owner_id" gorm:"primaryKey; not null"`
	Topic    Topic  `json:"topic" gorm:"primaryKey; not null"`
	Link     string `json:"link" gorm:"not null"`
}

func (r *Topic) UnmarshalJSON(data []byte) error {
	//unmarshal json data into string pointer
	v := new(string)
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	//convert string pointer to Type Role in val + check it match role
	val := Topic(*v)
	if val != Instagram && val != Facebook && val != Line && val != Linkedin && val != X && val != Whatsapp {
		return fmt.Errorf("%s is not a valid Role", val)
	}

	//update role
	*r = val
	return nil
}
