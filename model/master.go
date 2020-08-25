package model

// Title ...
type Title struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Gender ...
type Gender struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Location ...
type Location struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Department ...
type Department struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Status ...
type Status struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Product ...
type Product struct {
	ID              string `json:"id" bson:"id"`
	Name            string `json:"name" bson:"name"`
	RevenueCategory string `json:"revenue_category" bson:"revenue_category"`
}

// Partner ...
type Partner struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// Client ...
type Client struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Partner string `json:"partner" bson:"partner"`
}
