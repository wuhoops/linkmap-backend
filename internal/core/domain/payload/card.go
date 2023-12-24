package payload

type Card struct {
	CardId      string `json:"card_id"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type CardList struct {
	Card []Card `json:"card"`
}
