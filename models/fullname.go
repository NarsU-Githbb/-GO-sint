package models

// Mimics the FAKE mock_profiles.json schema

type MockProfile struct {
	SearchQuery string `json:"search_query"`
	Status      string `json:"status"`
	Person      struct {
		FullName        string `json:"full_name"`
		BirthYear       int    `json:"birth_year"`
		CurrentLocation string `json:"current_location"`
	} `json:"person"`
	ScrapedSources []ScrapedSource `json:"scraped_sources"`
}

type ScrapedSource struct {
	SourceType  string      `json:"source_type"`
	SourceName  string      `json:"source_name"`
	PhoneNumber string      `json:"phone_number,omitempty"`
	Address     MockAddress `json:"address,omitempty"`
}

type MockAddress struct {
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
}

type FullModel struct {
	FirstName   string
	LastName    string
	Address     MockAddress
	PhoneNumber string
}
