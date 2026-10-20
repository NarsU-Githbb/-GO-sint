package models

//Check the Json stuff later

//ip-api.com and ipapi.co for ip search? is in legal?

//http://ip-api.com/json/{query}

type IPModel struct {
	ISP        string `json:"isp"`
	Location   Cordinates
	DataSource string `json:"data_source"`
}

type Cordinates struct {
	Latitude  int `json:"lat"`
	Longitude int `json:"lon"`
}
