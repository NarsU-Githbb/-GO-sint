package models

//use github api for both username / fullname??

//Check the Json stuff later

type UserModel struct {
	SecondLife bool   `json:"second_life"`
	X          bool   `json:"x"`
	LinkedIn   bool   `json:"linked_in"`
	Instagram  bool   `json:"instagram"`
	Github     bool   `json:"github"`
	DataSource string `json:"data_source"`
}
