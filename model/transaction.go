package model

// Transaction ...
type Transaction struct {
	ID              string `json:"id" bson:"id"`
	ClientID        string `json:"client_id" bson:"client_id"`
	Client          string `json:"client" bson:"client"`
	Industry        string `json:"industry" bson:"industry"`
	Department      string `json:"department" bson:"department"`
	ProductID       string `json:"product_id" bson:"product_id"`
	Product         string `json:"product" bson:"product"`
	RevenueCategory string `json:"revenue_category" bson:"revenue_category"`
	Amount          string `json:"amount" bson:"amount"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
}
