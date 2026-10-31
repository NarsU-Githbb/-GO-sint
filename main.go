package main

import (
	"context"
	"digital_detective/scraping"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

var (
	fullSearch     = flag.String("n", "", "Performs a full-name search.")
	ipSearch       = flag.String("ip", "", "Performs a IP search.")
	usernameSearch = flag.String("un", "", "Performs a username search.")
	rawOutput      = flag.Bool("raw", false, "Outputs raw results.")
)

func main() {

	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "Warning: .env file not found; using system environment variables")
	}

	//context is needed so that the program doesn't secretly keep going after ctrl + C

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	flag.Usage = printHelp
	flag.Parse()

	activeFlags := 0
	if *fullSearch != "" {
		activeFlags++
	}
	if *ipSearch != "" {
		activeFlags++
	}
	if *usernameSearch != "" {
		activeFlags++
	}

	if activeFlags != 1 {
		printHelp()
		os.Exit(1)
	}

	var argument string
	fns := scraping.NewFnScraper()
	ips := scraping.NewIpScraper()
	uns := scraping.NewUnScraper()

	switch {

	case *fullSearch != "":
		argument = *fullSearch
		fmt.Printf("Starting Full Name Search for: %s\n", argument)
		result, err := fns.ExecuteFullSearch(ctx, argument)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Full-name search failed: %v\n", err)
			os.Exit(1)
		}

		if *rawOutput {
			makeRawFilename := createFileName(argument)
			rawFileName := fmt.Sprintf("raw_n_%s.txt", makeRawFilename)
			rawFileContent := CreateRawFullOutput(result.Raw, rawFileName)

			err = WriteOutput(rawFileName, rawFileContent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing raw results: %v\n", err)
			} else {
				fmt.Printf("Raw results written to file: %s\n", rawFileName)
			}

		} else {

			// Normal output

			makeFilename := createFileName(argument)
			fileName := fmt.Sprintf("n_%s.txt", makeFilename)
			fileContent := CreateFullOutput(argument, result.Summary, fileName)

			fmt.Println(fileContent)

			err = WriteOutput(fileName, fileContent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing results: %v\n", err)
			} else {
				fmt.Printf("Results written to file: %s\n", fileName)
			}
		}

	case *ipSearch != "":
		argument = *ipSearch
		fmt.Printf("Starting IP Search for: %s\n", argument)
		result, err := ips.ExecuteIPSearch(ctx, argument)
		if err != nil {
			fmt.Fprintf(os.Stderr, "IP search failed: %v\n", err)
			os.Exit(1)
		}

		if *rawOutput {

			makeRawFilename := createFileName(argument)
			rawFileName := fmt.Sprintf("raw_n_%s.txt", makeRawFilename)
			rawFileContent := CreateRawIPOutput(result.Raw, rawFileName)

			err = WriteOutput(rawFileName, rawFileContent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing raw results: %v\n", err)
			} else {
				fmt.Printf("Raw results written to file: %s\n", rawFileName)
			}

		} else {

			makeFilename := createFileName(argument)
			fileName := fmt.Sprintf("ip_%s.txt", makeFilename)
			fileContent := CreateIPOutput(argument, result.Summary, fileName)

			fmt.Println("--- TEST OUTPUT START ---")
			fmt.Printf("Target Filename: %s\n", fileName)
			fmt.Println(fileContent)
			fmt.Println("--- TEST OUTPUT END ---")

			err = WriteOutput(fileName, fileContent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing results: %v\n", err)
			} else {
				fmt.Printf("Results written to file: %s\n", fileName)
			}
		}

	case *usernameSearch != "":
		argument = *usernameSearch
		fmt.Printf("Starting Username Search for: %s\n", argument)
		result, err := uns.ExecuteUsernameSearch(ctx, argument)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Username search failed: %v\n", err)
			os.Exit(1)
		}

		makeFilename := createFileName(argument)
		fileName := fmt.Sprintf("un_%s.txt", makeFilename)

		fileContent := CreateUsernameOutput(argument, result, fileName)

		fmt.Println("--- TEST OUTPUT START ---")
		fmt.Printf("Target Filename: %s\n", fileName)
		fmt.Println(fileContent)
		fmt.Println("--- TEST OUTPUT END ---")

		err = WriteOutput(fileName, fileContent)
		if err != nil {
			fmt.Printf("Error writing results to file: %v\n", err)
		} else {
			fmt.Printf("Result successfully written to file: %s\n", fileName)
		}
	}

}

func printHelp() {

	const (
		reset     = "\033[0m"
		bold      = "\033[1m"
		italic    = "\033[3m"
		underline = "\033[4m"
		blue      = "\033[34m"
		green     = "\033[32m"
		yellow    = "\033[33m"
		red       = "\033[31m"
		magenta   = "\033[95m"
	)

	fmt.Println()
	fmt.Println(italic + bold + "Usage: go run . " + blue + " [flag] <input>" + reset)
	fmt.Println()

	fmt.Println(yellow + bold + "Flags:" + reset)
	fmt.Println()
	fmt.Println("  " + green + "-n" + reset + "      Performs a full-name search.")
	fmt.Println("  " + green + "-ip" + reset + "      Performs a IP search.")
	fmt.Println("  " + green + "-un" + reset + "      Performs a username search.")
	fmt.Println("  " + green + "-raw" + reset + "      Returns raw API responses (add before -n or -ip flag).")
	fmt.Println()

	fmt.Println()

	fmt.Println(yellow + bold + "Examples:" + reset)
	fmt.Println()
	fmt.Println(bold + "1. go run . " + blue + "-n 'John Doe'" + reset)
	fmt.Println()
	fmt.Println(bold + "First Name:" + reset + " John")
	fmt.Println(bold + "Last Name:" + reset + " Doe")
	fmt.Println(bold + "Address:" + reset + " Mango street 12 A 4, 00100 Helsinki")
	fmt.Println(bold + "Phone Number:" + reset + " +358 40 123 4567")
	fmt.Println(bold + "Result written to file:" + reset + " n_john_doe.txt")
	fmt.Println()
	fmt.Println(bold + "2. go run . " + blue + " -ip 8.8.8.1" + reset)
	fmt.Println()
	fmt.Println(bold + "IP:" + reset + " 8.8.8.1")
	fmt.Println(bold + "ISP:" + reset + " Google LLC")
	fmt.Println(bold + "Location:" + reset + " 37.422000,-122.085000")
	fmt.Println(bold + "Validated:" + reset + " true")
	fmt.Println(bold + "Warnings:" + reset + " None")
	fmt.Println(bold + "Abuse Confidence Score:" + reset + " Not available")
	fmt.Println(bold + "Total Reports:" + reset + " Not available")
	fmt.Println(bold + "Last Reported At:" + reset + " Not searched")
	fmt.Println(bold + "Result written to file:" + reset + " ip_8_8_8_1.txt")
	fmt.Println()
	fmt.Println(bold + "3. go run . " + blue + "-un torvalds" + reset)
	fmt.Println()
	fmt.Println(bold + "X:" + reset + " true")
	fmt.Println(bold + "Youtube:" + reset + " false")
	fmt.Println(bold + "Soundcloud:" + reset + " true")
	fmt.Println(bold + "Github:" + reset + " true")
	fmt.Println(bold + "Gitea:" + reset + " true")
	fmt.Println(bold + "Result written to file:" + reset + " un_torvalds.txt")
	fmt.Println()
}
