package payload

import "backend/internal/core/domain/database"

type Social struct {
	SocialId string         `json:"card_id"`
	OwnerId  string         `json:"owner_id"`
	Topic    database.Topic `json:"topic"`
	Link     string         `json:"link"`
}
