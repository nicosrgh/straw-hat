package model

// Employee ...
type Employee struct {
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
