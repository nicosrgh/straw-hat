package model

// Employee ...
type Employee struct {
	ID           string `json:"id" bson:"id"`
	NIP          string `json:"nip" bson:"nip"`
	Fullname     string `json:"fullname" bson:"fullname"`
	StatusID     string `json:"status_id" bson:"status_id"`
	Status       string `json:"status" bson:"status"`
	GenderID     string `json:"gender_id" bson:"gender_id"`
	Gender       string `json:"gender" bson:"gender"`
	DepartmentID string `json:"department_id" bson:"department_id"`
	Department   string `json:"department" bson:"department"`
	LocationID   string `json:"location_id" bson:"location_id"`
	Location     string `json:"location" bson:"location"`
	TitleID      string `json:"title_id" bson:"title_id"`
	Title        string `json:"title" bson:"title"`
	Birthdate    string `json:"birthdate" bson:"birthdate"`
	JoinDate     string `json:"join_date" bson:"join_date"`
}

// SourceEmployee ...
type SourceEmployee struct {
	ID                string `json:"id" bson:"id"`
	SourceID          string `json:"source_id" bson:"source_id"`
	NIP               string `json:"nip" bson:"nip"`
	Fullname          string `json:"fullname" bson:"fullname"`
	Status            string `json:"status" bson:"status"`
	Gender            string `json:"gender" bson:"gender"`
	Department        string `json:"department" bson:"department"`
	Location          string `json:"location" bson:"location"`
	Title             string `json:"title" bson:"title"`
	Birthdate         string `json:"birthdate" bson:"birthdate"`
	JoinDate          string `json:"join_date" bson:"join_date"`
	Day               string `json:"day" bson:"day"`
	Week              string `json:"week" bson:"week"`
	Month             string `json:"month" bson:"month"`
	Year              string `json:"year" bson:"year"`
	TotalEmployee     string `json:"total_employee" bson:"total_employee"`
	MaleEmployee      string `json:"male_employee" bson:"male_employee"`
	FemaleEmployee    string `json:"female_employee" bson:"female_employee"`
	FullTimeEmployee  string `json:"full_time_employee" bson:"full_time_employee"`
	ProbationEmployee string `json:"probation_employee" bson:"probation_employee"`
}

// DatamartTitle ...
type DatamartTitle struct {
	ID        string `json:"id" bson:"id"`
	TitleFact int    `json:"title_fact" bson:"title_fact"`
}

// DatamartEmployeeTotal ...
type DatamartEmployeeTotal struct {
	ID                string `json:"id" bson:"id"`
	TotalEmployee     string `json:"total_employee" bson:"total_employee"`
	MaleEmployee      string `json:"male_employee" bson:"male_employee"`
	FemaleEmployee    string `json:"female_employee" bson:"female_employee"`
	FullTimeEmployee  string `json:"full_time_employee" bson:"full_time_employee"`
	ProbationEmployee string `json:"probation_employee" bson:"probation_employee"`
}
