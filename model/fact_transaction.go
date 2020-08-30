package model

// FactTransactionClient ...
type FactTransactionClient struct {
	Count     string `json:"count" bson:"count"`
	ClientID  string `json:"client_id" bson:"client_id"`
	Client    string `json:"client" bson:"client"`
	CreatedAt string `json:"created_date" bson:"created_date"`
	Day       string `json:"day" bson:"day"`
	Month     string `json:"month" bson:"month"`
	Year      string `json:"year" bson:"year"`
}

// FactTransactionProduct ...
type FactTransactionProduct struct {
	Count     string `json:"count" bson:"count"`
	ProductID string `json:"product_id" bson:"product_id"`
	Product   string `json:"product" bson:"product"`
	CreatedAt string `json:"created_date" bson:"created_date"`
	Day       string `json:"day" bson:"day"`
	Month     string `json:"month" bson:"month"`
	Year      string `json:"year" bson:"year"`
}

// FactTransactionPartner ...
type FactTransactionPartner struct {
	Count     string `json:"count" bson:"count"`
	PartnerID string `json:"partner_id" bson:"partner_id"`
	Partner   string `json:"partner" bson:"partner"`
	CreatedAt string `json:"created_date" bson:"created_date"`
	Day       string `json:"day" bson:"day"`
	Month     string `json:"month" bson:"month"`
	Year      string `json:"year" bson:"year"`
}
