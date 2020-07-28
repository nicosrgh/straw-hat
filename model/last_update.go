package model

// LastUpdate ...
type LastUpdate struct {
	ID        string `json:"id" bson:"id"`
	Action    string `json:"action" bson:"action"`
	LastID    string `json:"last_id" bson:"last_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
