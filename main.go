package main

import (
	"context"
	"digital_detective/scraping"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	fullSearch     = flag.String("n", "", "Performs a full-name search.")
	ipSearch       = flag.String("ip", "", "Performs a IP search.")
	usernameSearch = flag.String("un", "", "Performs a username search.")
)

func main() {

	//context is needed so that the program doesn't secretly keep going even after ctrl + C

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

	if activeFlags != 1 { //allow just one for now, but change if decide to do multi-input
		printHelp()
		os.Exit(1)
	}

	var argument string

	scraper := scraping.NewScraper()

	switch {

	case *fullSearch != "":
		argument = *fullSearch
		fmt.Printf("Starting Full Name Search for: %s\n", argument)
		//scraper.ExecuteNameSearch(ctx, argument)

	case *ipSearch != "":
		argument = *ipSearch
		fmt.Printf("Starting IP Search for: %s\n", argument)
		//scraper.ExecuteIPSearch(ctx, argument)

	case *usernameSearch != "":
		argument = *usernameSearch
		fmt.Printf("Starting Username Search for: %s\n", argument)
		result := scraper.ExecuteUsernameSearch(ctx, argument)

		fileName := fmt.Sprintf("un_%s.txt", argument)
		fileContent := CreateUsernameOutput(argument, result, fileName)

		err := WriteOutput(fileName, fileContent)
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
	fmt.Println(italic + bold + "Usage: go run data_digger.go" + blue + " [options] <input>" + reset)
	fmt.Println()

	fmt.Println(yellow + bold + "Flags:" + reset)
	fmt.Println()
	fmt.Println("  " + green + "-n" + reset + "      Performs a full-name search.")
	fmt.Println("  " + green + "-ip" + reset + "      Performs a IP search.")
	fmt.Println("  " + green + "-un" + reset + "      Performs a username search.")
	fmt.Println()

	fmt.Println()

	fmt.Println(yellow + bold + "Examples:" + reset)
	fmt.Println()
	fmt.Println(bold + "1. go run data_digger.go " + blue + "-n Dwight Schrute" + reset)
	fmt.Println()
	fmt.Println(bold + "First Name:" + reset + " Dwight")
	fmt.Println(bold + "Last Name:" + reset + " Schurute")
	fmt.Println(bold + "Address:" + reset + " Schrute Farms, Rural Route 6, Honesdale, PA 18431")
	fmt.Println(bold + "Phone Number:" + reset + " 308-826-5962")
	fmt.Println(bold + "Result written to file:" + reset + "n_dwight_schrute.txt")
	fmt.Println()
	fmt.Println(bold + "2. go run data_digger.go" + blue + " -ip 192.168.1.1" + reset)
	fmt.Println()
	fmt.Println(bold + "ISP:" + reset + " Dunder Mifflin")
	fmt.Println(bold + "Location:" + reset + " 34.210766, -118.436440")
	fmt.Println(bold + "Result written to file:" + reset + " ip_192_168_1_1.txt")
	fmt.Println()
	fmt.Println(bold + "3. go run data_digger.go " + blue + "-un @recyclops" + reset)
	fmt.Println()
	fmt.Println(bold + "Second Life:" + reset + " yes")
	fmt.Println(bold + "X:" + reset + " no")
	fmt.Println(bold + "LinkedIn:" + reset + " no")
	fmt.Println(bold + "Instagram:" + reset + " no")
	fmt.Println(bold + "Github:" + reset + " no")
	fmt.Println(bold + "Result written to file:" + reset + " un_recyclops.txt")
	fmt.Println()
}
