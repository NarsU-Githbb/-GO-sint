package main

import (
	"digital_detective/models"
	"fmt"
	"os"
	"strings"
)

func createFileName(argument string) string {

	argument = strings.ToLower(strings.TrimSpace(argument))

	var b strings.Builder
	lastUnderscore := false

	for _, r := range argument {
		switch {
		case r >= 'a' && r <= 'z',
			r >= '0' && r <= '9',
			r == '-', r == '_':
			b.WriteRune(r)
			lastUnderscore = false

		case r == ' ' || r == '.':
			if b.Len() > 0 && !lastUnderscore {
				b.WriteByte('_')
				lastUnderscore = true
			}

		default:

		}
	}

	filenameFix := strings.Trim(b.String(), "_")
	return filenameFix

}

func CreateFullOutput(fullname string, result models.FullModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for %s ---\n", fullname)
	output += fmt.Sprintf("First Name: %s\n", result.FirstName)
	output += fmt.Sprintf("Last Name: %s\n", result.LastName)
	output += fmt.Sprintf(
		"Address: %s, %s, %s\n",
		result.Address.Street,
		result.Address.PostalCode,
		result.Address.City,
	)
	output += fmt.Sprintf("Phone Number: %s\n", result.PhoneNumber)
	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func CreateIPOutput(IP string, result models.IPModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", IP)
	output += fmt.Sprintf("IP: %s\n", result.IP)
	output += fmt.Sprintf("ISP: %s\n", result.ISP)
	output += fmt.Sprintf("Location: %.6f,%.6f\n",
		result.Location.Latitude,
		result.Location.Longitude)
	output += fmt.Sprintf("Validated: %t\n", result.IsValidated)
	output += "Warnings:\n"
	if len(result.Warning) == 0 {
		output += "  None\n"
	} else {
		for _, war := range result.Warning {
			output += fmt.Sprintf("  - %s\n", war)
		}
	}

	// AbuseIPDB fields are optional.

	if result.AbuseConfidenceScore != nil {
		output += fmt.Sprintf("Abuse Confidence Score: %d\n", *result.AbuseConfidenceScore)
	} else {
		output += "Abuse Confidence Score: Not available\n"
	}

	if result.TotalReports != nil {
		output += fmt.Sprintf("Total Reports: %d\n", *result.TotalReports)
	} else {
		output += "Total Reports: Not available\n"
	}

	if result.LastReportedAt != "" {
		output += fmt.Sprintf("Last Reported At: %s\n", result.LastReportedAt)
	} else {
		output += "Last Reported At: Not available\n"
	}

	output += "Errors:\n"
	if len(result.Errors) == 0 {
		output += "  None\n"
	} else {
		for _, er := range result.Errors {
			output += fmt.Sprintf("  - %s\n", er)
		}
	}

	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func CreateUsernameOutput(username string, result models.UserModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", username)
	output += fmt.Sprintf("X:           %t\n", result.X)
	output += fmt.Sprintf("Youtube:    %t\n", result.Youtube)
	output += fmt.Sprintf("Soundcloud: %t\n", result.Soundcloud)
	output += fmt.Sprintf("Github:      %t\n", result.Github)
	output += fmt.Sprintf("Gitea:   %t\n", result.Gitea)
	output += "Errors:\n"
	if len(result.Errors) == 0 {
		output += "  None\n"
	} else {
		for _, er := range result.Errors {
			output += fmt.Sprintf("  - %s\n", er)
		}
	}
	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func WriteOutput(outputPath string, output string) error {

	pathName := strings.TrimSuffix(outputPath, ".txt")

	counter := 1
	finalPath := outputPath

	for {

		_, err := os.Stat(finalPath)
		if os.IsNotExist(err) {
			break
		}

		if err != nil {
			return err // if different filesystem error occurred
		}
		finalPath = fmt.Sprintf("%s(%d).txt", pathName, counter)
		counter++
	}

	outputFile, err := os.OpenFile(finalPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	_, err = outputFile.Write([]byte(output))
	if err != nil {
		return err
	}

	return nil
}
