package main

import (
	"digital_detective/models"
	"fmt"
	"os"
	"path/filepath"
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

func statusIcon(ok bool) string {
	if ok {
		return "[+]"
	}
	return "[-]"
}

func asciiBar(value, max, width int) string {
	if max <= 0 || width <= 0 {
		return "[" + strings.Repeat(".", maxInt(width, 0)) + "]"
	}

	if value < 0 {
		value = 0
	}
	if value > max {
		value = max
	}

	filled := value * width / max

	return "[" +
		strings.Repeat("#", filled) +
		strings.Repeat(".", width-filled) +
		"]"
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CreateFullOutput(fullname string, result models.FullModel, filename string) string {

	var output strings.Builder

	output.WriteString("============================================\n")
	output.WriteString("           (G)OSINT CLI TOOL              \n")
	output.WriteString("                PERSON SCAN                \n")
	output.WriteString("============================================\n\n")

	fmt.Fprintf(&output, "  Target: %s\n\n", fullname)

	output.WriteString("  PROFILE DETAILS\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(&output, "  First Name: %s\n", result.FirstName)
	fmt.Fprintf(&output, "  Last Name:  %s\n", result.LastName)

	output.WriteString("\n  ADDRESS\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(&output, "  Street: %s\n", result.Address.Street)
	fmt.Fprintf(&output, "  Postal: %s\n", result.Address.PostalCode)
	fmt.Fprintf(&output, "  City:   %s\n", result.Address.City)

	output.WriteString("\n  CONTACT\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(&output, "  Phone: %s\n", result.PhoneNumber)

	output.WriteString("\n============================================\n")
	fmt.Fprintf(&output, "  Output: %s\n", filename)
	output.WriteString("============================================\n\n")

	return output.String()
}

func CreateIPOutput(IP string, result models.IPModel, filename string) string {

	var output strings.Builder

	output.WriteString("============================================\n")
	output.WriteString("           (G)OSINT CLI TOOL               \n")
	output.WriteString("                 IP SCAN                   \n")
	output.WriteString("============================================\n\n")

	// Target information
	output.WriteString("  TARGET INFORMATION\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(&output, "  Target:      %s\n", IP)
	fmt.Fprintf(&output, "  IP Address:  %s\n", result.IP)
	fmt.Fprintf(&output, "  ISP:         %s\n", result.ISP)
	fmt.Fprintf(
		&output,
		"  Coordinates: %.6f, %.6f\n",
		result.Location.Latitude,
		result.Location.Longitude,
	)

	// Validation
	output.WriteString("\n  VALIDATION\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(
		&output,
		"  Status: %s\n",
		map[bool]string{
			true:  "[+] Validated",
			false: "[-] Not validated",
		}[result.IsValidated],
	)

	// Abuse assessment
	output.WriteString("\n  ABUSE ASSESSMENT\n")
	output.WriteString("  ------------------------------------------\n")

	if result.AbuseConfidenceScore != nil {
		score := *result.AbuseConfidenceScore

		fmt.Fprintf(
			&output,
			"  Confidence: %s %d/100\n",
			asciiBar(score, 100, 20),
			score,
		)
	} else {
		output.WriteString("  Confidence: Not available\n")
	}

	if result.TotalReports != nil {
		fmt.Fprintf(&output, "  Total Reports: %d\n", *result.TotalReports)
	} else {
		output.WriteString("  Total Reports: Not available\n")
	}

	if result.LastReportedAt != "" {
		fmt.Fprintf(
			&output,
			"  Last Reported: %s\n",
			result.LastReportedAt,
		)
	} else {
		output.WriteString("  Last Reported: Not available\n")
	}

	// Warnings
	output.WriteString("\n  WARNINGS\n")
	output.WriteString("  ------------------------------------------\n")

	if len(result.Warning) == 0 {
		output.WriteString("  [+] No warnings\n")
	} else {
		for _, warning := range result.Warning {
			fmt.Fprintf(&output, "  [!] %s\n", warning)
		}
	}

	// Errors
	output.WriteString("\n  ERRORS\n")
	output.WriteString("  ------------------------------------------\n")

	if len(result.Errors) == 0 {
		output.WriteString("  [+] No errors\n")
	} else {
		for _, errMsg := range result.Errors {
			fmt.Fprintf(&output, "  [-] %s\n", errMsg)
		}
	}

	// Footer
	output.WriteString("\n============================================\n")
	fmt.Fprintf(&output, "  Output: %s\n", filename)
	output.WriteString("============================================\n\n")

	return output.String()
}

func CreateUsernameOutput(username string, result models.UserModel, filename string) string {

	platforms := []struct {
		name  string
		found bool
	}{
		{"X", result.X},
		{"YouTube", result.Youtube},
		{"SoundCloud", result.Soundcloud},
		{"GitHub", result.Github},
		{"Gitea", result.Gitea},
	}

	found := 0
	var output strings.Builder

	output.WriteString("============================================\n")
	output.WriteString("           (G)OSINT CLI TOOL               \n")
	output.WriteString("              USERNAME SCAN                \n")
	output.WriteString("============================================\n\n")

	fmt.Fprintf(&output, "  Target: @%s\n\n", username)
	output.WriteString("  PLATFORM PRESENCE\n")
	output.WriteString("  ------------------------------------------\n")

	for _, platform := range platforms {
		label := "NOT FOUND"
		if platform.found {
			label = "FOUND"
			found++
		}

		fmt.Fprintf(
			&output,
			"  %-12s %-3s %s\n",
			platform.name,
			statusIcon(platform.found),
			label,
		)
	}

	total := len(platforms)
	percentage := 0
	if total > 0 {
		percentage = found * 100 / total
	}

	output.WriteString("\n  SCAN SUMMARY\n")
	output.WriteString("  ------------------------------------------\n")
	fmt.Fprintf(
		&output,
		"  Found:    %s %d/%d\n",
		asciiBar(found, total, 20),
		found,
		total,
	)
	fmt.Fprintf(&output, "  Coverage: %d%%\n", percentage)

	output.WriteString("\n  ERRORS\n")
	output.WriteString("  ------------------------------------------\n")

	if len(result.Errors) == 0 {
		output.WriteString("  [+] No errors\n")
	} else {
		for _, errMsg := range result.Errors {
			fmt.Fprintf(&output, "  [-] %s\n", errMsg)
		}
	}

	output.WriteString("\n============================================\n")
	fmt.Fprintf(&output, "  Output: %s\n", filename)
	output.WriteString("============================================\n\n")

	return output.String()
}

func WriteOutput(outputPath string, output string) error {

	if err := os.MkdirAll("reports", 0755); err != nil {
		return err
	}

	// Extract filename without extension
	filename := filepath.Base(outputPath)
	pathName := strings.TrimSuffix(filename, ".txt")

	// Construct initial output path
	finalPath := filepath.Join("reports", filename)

	counter := 1

	for {

		_, err := os.Stat(finalPath)
		if os.IsNotExist(err) {
			break
		}

		if err != nil {
			return err // if different filesystem error occurred
		}
		finalPath = filepath.Join(
			"reports",
			fmt.Sprintf("%s(%d).txt", pathName, counter),
		)

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
