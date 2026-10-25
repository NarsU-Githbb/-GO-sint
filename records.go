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

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", fullname)
	output += fmt.Sprintf("First Name: %s\n", result.FirstName)
	output += fmt.Sprintf("Last Name: %s\n", result.LastName)
	output += fmt.Sprintf("Adress: %v\n", result.Address)
	output += fmt.Sprintf("Phone Number: %s\n", result.PhoneNumber)
	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func CreateIPOutput(IP string, result models.IPModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", IP)
	output += fmt.Sprintf("IP: %s\n", result.IP)
	output += fmt.Sprintf("ISP: %s\n", result.ISP)
	output += fmt.Sprintf("Location: %v\n", result.Location)
	output += fmt.Sprintf("Validated: %t\n", result.IsValidated)
	output += fmt.Sprintf("Warnings: %v\n", result.Warning)
	output += fmt.Sprintf("Abuse Confidence Score: %d\n", *result.AbuseConfidenceScore)
	output += fmt.Sprintf("Total Reports: %d\n", *result.TotalReports)
	output += fmt.Sprintf("Last Recorded At: %s\n", result.LastReportedAt)
	output += fmt.Sprintf("Errors: %v\n", result.Errors)
	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func CreateUsernameOutput(username string, result models.UserModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", username)
	output += fmt.Sprintf("Second Life: %t\n", result.SecondLife)
	output += fmt.Sprintf("X:           %t\n", result.X)
	output += fmt.Sprintf("LinkedIn:    %t\n", result.LinkedIn)
	output += fmt.Sprintf("Instagram:   %t\n", result.Instagram)
	output += fmt.Sprintf("Github:      %t\n", result.Github)
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
