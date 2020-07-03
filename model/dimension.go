package model

// DimensionGender ...
type DimensionGender struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// DimensionTitle ...
type DimensionTitle struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
