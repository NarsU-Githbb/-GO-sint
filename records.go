package main

import (
	"digital_detective/models"
	"fmt"
	"os"
)

func CreateFullOutput(fullname string, result models.FullModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", fullname)
	output += fmt.Sprintf("First Name: %t\n", result.FirstName)
	output += fmt.Sprintf("Last Name: %t\n", result.LastName)
	output += fmt.Sprintf("Adress: %t\n", result.Address)
	output += fmt.Sprintf("Phone Number: %t\n", result.PhoneNumber)
	output += fmt.Sprintf("Result written to file:      %s\n\n", filename)

	return output
}

func CreateIPOutput(IP string, result models.IPModel, filename string) string {

	output := fmt.Sprintf("--- Scan Results for @%s ---\n", IP)
	output += fmt.Sprintf("ISP: %t\n", result.ISP)
	output += fmt.Sprintf("Location: %t\n", result.Location)
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

	outputFile, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	_, err = outputFile.Write([]byte(output))
	if err != nil {
		return err
	}
	defer outputFile.Close()
	return nil
}
