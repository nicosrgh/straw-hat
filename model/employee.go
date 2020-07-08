package model

// Employee ...
type Employee struct {
	ID         string `json:"id" bson:"id"`
	NIP        string `json:"nip" bson:"nip"`
	Fullname   string `json:"fullname" bson:"fullname"`
	Status     string `json:"status" bson:"status"`
	Gender     string `json:"gender" bson:"gender"`
	Department string `json:"department" bson:"department"`
	Location   string `json:"location" bson:"location"`
	Title      string `json:"title" bson:"title"`
	Birthdate  string `json:"birthdate" bson:"birthdate"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	Count      string `json:"count" bson:"count"`
}

// SourceEmployee ...
type SourceEmployee struct {
	ID         string `json:"id" bson:"id"`
	SourceID   string `json:"source_id" bson:"source_id"`
	NIP        string `json:"nip" bson:"nip"`
	Fullname   string `json:"fullname" bson:"fullname"`
	Status     string `json:"status" bson:"status"`
	Gender     string `json:"gender" bson:"gender"`
	Department string `json:"department" bson:"department"`
	Location   string `json:"location" bson:"location"`
	Title      string `json:"title" bson:"title"`
	Birthdate  string `json:"birthdate" bson:"birthdate"`
	JoinDate   string `json:"join_date" bson:"join_date"`
}

// DatamartTitle ...
type DatamartTitle struct {
	ID        string `json:"id" bson:"id"`
	TitleFact int    `json:"title_fact" bson:"title_fact"`
}
