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

// DimensionLocation ...
type DimensionLocation struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// DimensionTime ...
type DimensionTime struct {
	ID          string `json:"id" bson:"id"`
	DBDate      string `json:"db_date" bson:"db_date"`
	Year        string `json:"year" bson:"year"`
	Month       string `json:"month" bson:"month"`
	Day         string `json:"day" bson:"day"`
	Quarter     string `json:"quarter" bson:"quarter"`
	Week        string `json:"week" bson:"week"`
	DayName     string `json:"day_name" bson:"day_name"`
	MonthName   string `json:"month_name" bson:"month_name"`
	HolidayFlag string `json:"holiday_flag" bson:"holiday_flag"`
	WeekendFlag string `json:"weekend_flag" bson:"weekend_flag"`
	Event       string `json:"event" bson:"event"`
}
