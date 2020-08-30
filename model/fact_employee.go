package model

// FactGenderEmployee ...
type FactGenderEmployee struct {
	Count      string `json:"count" bson:"count"`
	GenderID   string `json:"gender_id" bson:"gender_id"`
	Gender     string `json:"gender" bson:"gender"`
	LocationID string `json:"location_id" bson:"location_id"`
	Location   string `json:"location" bson:"location"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	Day        string `json:"day" bson:"day"`
	Month      string `json:"month" bson:"month"`
	Year       string `json:"year" bson:"year"`
}

// FactLocationEmployee ...
type FactLocationEmployee struct {
	Count      string `json:"count" bson:"count"`
	LocationID string `json:"location_id" bson:"location_id"`
	Location   string `json:"location" bson:"location"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	Day        string `json:"day" bson:"day"`
	Month      string `json:"month" bson:"month"`
	Year       string `json:"year" bson:"year"`
}

// FactStatusEmployee ...
type FactStatusEmployee struct {
	Count      string `json:"count" bson:"count"`
	StatusID   string `json:"status_id" bson:"status_id"`
	Status     string `json:"status" bson:"status"`
	LocationID string `json:"location_id" bson:"location_id"`
	Location   string `json:"location" bson:"location"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	Day        string `json:"day" bson:"day"`
	Month      string `json:"month" bson:"month"`
	Year       string `json:"year" bson:"year"`
}
