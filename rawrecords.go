package main

import (
	"digital_detective/models"
	"fmt"
	"strings"
)

func CreateRawFullOutput(result models.MockProfile, filename string) string {

	var output strings.Builder

	fmt.Fprintf(&output, "--- Raw Scan Results ---\n")
	fmt.Fprintf(&output, "Search Query: %s\n", result.SearchQuery)
	fmt.Fprintf(&output, "Status: %s\n", result.Status)

	// Person information
	fmt.Fprintf(&output, "\n--- Person Information ---\n")
	fmt.Fprintf(&output, "Full Name: %s\n", result.Person.FullName)
	fmt.Fprintf(&output, "Birth Year: %d\n", result.Person.BirthYear)
	fmt.Fprintf(&output, "Current Location: %s\n", result.Person.CurrentLocation)

	// Scraped sources
	fmt.Fprintf(&output, "\n--- Scraped Sources ---\n")

	if len(result.ScrapedSources) == 0 {
		fmt.Fprintf(&output, "No scraped sources found.\n")
	}

	for i, source := range result.ScrapedSources {
		fmt.Fprintf(&output, "\nSource #%d\n", i+1)
		fmt.Fprintf(&output, "Source Type: %s\n", source.SourceType)
		fmt.Fprintf(&output, "Source Name: %s\n", source.SourceName)
		fmt.Fprintf(&output, "Phone Number: %s\n", source.PhoneNumber)

		fmt.Fprintf(&output, "Address:\n")
		fmt.Fprintf(&output, "  Street: %s\n", source.Address.Street)
		fmt.Fprintf(&output, "  Postal Code: %s\n", source.Address.PostalCode)
		fmt.Fprintf(&output, "  City: %s\n", source.Address.City)
	}

	fmt.Fprintf(&output, "\nResult written to file: %s\n", filename)

	return output.String()
}

func CreateRawIPOutput(result models.RawIPResults, filename string) string {

	var output strings.Builder

	fmt.Fprintf(&output, "--- Raw IP Scan Results ---\n")

	// IP-API response
	fmt.Fprintf(&output, "\n--- IP-API Response ---\n")

	if result.IPAPI != nil {
		fmt.Fprintf(&output, "Status: %s\n", result.IPAPI.Status)
		fmt.Fprintf(&output, "Message: %s\n", result.IPAPI.Message)
		fmt.Fprintf(&output, "Query: %s\n", result.IPAPI.Query)
		fmt.Fprintf(&output, "ISP: %s\n", result.IPAPI.ISP)
		fmt.Fprintf(&output, "City: %s\n", result.IPAPI.City)
		fmt.Fprintf(&output, "Country: %s\n", result.IPAPI.Country)
		fmt.Fprintf(&output, "Latitude: %f\n", result.IPAPI.Lat)
		fmt.Fprintf(&output, "Longitude: %f\n", result.IPAPI.Lon)
		fmt.Fprintf(&output, "AS: %s\n", result.IPAPI.AS)
	} else {
		fmt.Fprintf(&output, "No IP-API response available.\n")
	}

	// IPWhoIs response
	fmt.Fprintf(&output, "\n--- IPWhoIs Response ---\n")

	if result.IPWhoIs != nil {
		fmt.Fprintf(&output, "Success: %t\n", result.IPWhoIs.Success)
		fmt.Fprintf(&output, "IP: %s\n", result.IPWhoIs.IP)
		fmt.Fprintf(&output, "ISP: %s\n", result.IPWhoIs.ISP)
		fmt.Fprintf(&output, "City: %s\n", result.IPWhoIs.City)
		fmt.Fprintf(&output, "Country: %s\n", result.IPWhoIs.Country)
		fmt.Fprintf(&output, "Latitude: %f\n", result.IPWhoIs.Lat)
		fmt.Fprintf(&output, "Longitude: %f\n", result.IPWhoIs.Lon)
		fmt.Fprintf(&output, "ASN: %s\n", result.IPWhoIs.ASN)
	} else {
		fmt.Fprintf(&output, "No IPWhoIs response available.\n")
	}

	// AbuseIPDB response
	fmt.Fprintf(&output, "\n--- AbuseIPDB Response ---\n")

	if result.AbuseIP != nil {
		fmt.Fprintf(&output, "IP Address: %s\n", result.AbuseIP.IPAddress)
		fmt.Fprintf(&output, "Public IP: %t\n", result.AbuseIP.IsPublic)
		fmt.Fprintf(&output, "Country Code: %s\n", result.AbuseIP.CountryCode)
		fmt.Fprintf(&output, "ISP: %s\n", result.AbuseIP.ISP)

		fmt.Fprintf(&output, "Abuse Confidence Score: ")
		if result.AbuseIP.AbuseConfidenceScore != nil {
			fmt.Fprintf(&output, "%d\n", *result.AbuseIP.AbuseConfidenceScore)
		} else {
			fmt.Fprintf(&output, "Not available\n")
		}

		fmt.Fprintf(&output, "Total Reports: ")
		if result.AbuseIP.TotalReports != nil {
			fmt.Fprintf(&output, "%d\n", *result.AbuseIP.TotalReports)
		} else {
			fmt.Fprintf(&output, "Not available\n")
		}

		fmt.Fprintf(&output, "Last Reported At: %s\n", result.AbuseIP.LastReportedAt)
		fmt.Fprintf(&output, "ASN: %d\n", result.AbuseIP.ASN)
	} else {
		fmt.Fprintf(&output, "No AbuseIPDB response available.\n")
	}

	fmt.Fprintf(&output, "\nResult written to file: %s\n", filename)

	return output.String()
}
