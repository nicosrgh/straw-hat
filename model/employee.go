package model

// Employee ...
type Employee struct {
	ID         int32  `json:"id" bson:"id"`
	NIP        string `json:"nip" bson:"nip"`
	Fullname   string `json:"fullname" bson:"fullname"`
	Status     string `json:"status" bson:"status"`
	Gender     string `json:"gender" bson:"gender"`
	Department string `json:"department" bson:"department"`
	Title      string `json:"title" bson:"title"`
	Birthdate  string `json:"birthdate" bson:"birthdate"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	ResignDate string `json:"resign_date" bson:"resign_date"`
}

// SourceEmployee ...
type SourceEmployee struct {
	ID         int32  `json:"id" bson:"id"`
	SourceID   int32  `json:"source_id" bson:"source_id"`
	NIP        string `json:"nip" bson:"nip"`
	Fullname   string `json:"fullname" bson:"fullname"`
	Status     string `json:"status" bson:"status"`
	Gender     string `json:"gender" bson:"gender"`
	Department string `json:"department" bson:"department"`
	Title      string `json:"title" bson:"title"`
	Birthdate  string `json:"birthdate" bson:"birthdate"`
	JoinDate   string `json:"join_date" bson:"join_date"`
	ResignDate string `json:"resign_date" bson:"resign_date"`
}
