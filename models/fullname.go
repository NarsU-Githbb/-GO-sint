package models

//Check the Json stuff later

type FullModel struct {
	FirstName   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Adress      string `json:"adress"`
	PhoneNumber string `json:"phone_number"`
	DataSource  string `json:"data_source"`
}
