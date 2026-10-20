package main

import (
	"digital_detective/models"
	"fmt"
	"os"
)

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
