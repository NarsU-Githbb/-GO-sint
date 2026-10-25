package models

type IPModel struct {
	IP          string
	ISP         string
	Location    Coordinates
	IsValidated bool     // True if APIs agree on core metrics
	Warning     []string // Populated if the APIs disagree

	// AbuseIPDB threat intelligence only
	AbuseConfidenceScore *int   `json:"abuse_confidence_score,omitempty"`
	TotalReports         *int   `json:"total_reports,omitempty"`
	LastReportedAt       string `json:"last_reported_at,omitempty"`
	Errors               []string
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// Target schema for API 1: http://ip-api.com/json/{ip}

type IpApiResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message,omitempty"`
	Query   string  `json:"query"`
	ISP     string  `json:"isp"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	AS      string  `json:"as,omitempty"`
}

// Target schema for API 2: https://ipwhois.app/json/{ip}

type IpWhoIsResponse struct {
	Success bool    `json:"success"`
	IP      string  `json:"ip"`
	ISP     string  `json:"isp"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
	ASN     string  `json:"asn,omitempty"`
}

// Target schema for API 3: https://api.abuseipdb.com/api/v2//check?ipAddress={ip}

// IPAbuseResponse represents the outer AbuseIPDB response.
type IPAbuseResponse struct {
	Data IPAbuseReport `json:"data"`
}

// IPAbuseReport represents AbuseIPDB's nested data object.
type IPAbuseReport struct {
	IPAddress            string `json:"ipAddress"`
	IsPublic             bool   `json:"isPublic"`
	CountryCode          string `json:"countryCode,omitempty"`
	ISP                  string `json:"isp,omitempty"`
	AbuseConfidenceScore *int   `json:"abuseConfidenceScore,omitempty"`
	TotalReports         *int   `json:"totalReports,omitempty"`
	LastReportedAt       string `json:"lastReportedAt,omitempty"`
	ASN                  int    `json:"asn,omitempty"`
}
