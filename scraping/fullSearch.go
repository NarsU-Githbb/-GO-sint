package scraping

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

//"private_ditective/models"

// leave as an extra: if I want to add investigation links in to the report

func CreateFullUrl(fullName string) string {

	// Wrap the name in quotes
	quotedName := fmt.Sprintf(`"%s"`, fullName)
	searchQuery := quotedName + " in:fullname"

	// url.QueryEscape encodes the entire string correctly for ?q=
	encodedQuery := url.QueryEscape(searchQuery)

	URL := fmt.Sprintf("https://github.com/search?q=%s", encodedQuery)
	fullURL := URL + "&type=users"

	fmt.Println(fullURL)
	return fullURL
}

//do this last as it's the hardest category

//it's actually Illegal to search addresses & phone numbers of private individuals

func ExecuteNameSearch(fullName string) string {

	safeName := url.QueryEscape(fullName)
	targetURL := "https://github.com" + safeName

	resp, err := http.Get(targetURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}
