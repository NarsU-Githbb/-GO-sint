package models

//Check the Json stuff later

//ip-api.com and ipapi.co for ip search? is in legal?

type IPModel struct {
	ISP        string `json:"isp"`
	Location   Cordinates
	DataSource string `json:"data_source"`
}

type Cordinates struct {
	Vertical   int `json:"vertical"`
	Horizontal int `json:"horizontal"`
}
