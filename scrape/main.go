package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var url = "https://www.w3.org/Style/CSS/all-properties.en.html"

func main() {
	strs, err := Scrape()

	if err != nil {
		log.Fatalf("Error scraping CSS properties: %v", err)
		return
	}

	fmt.Printf("Found %d CSS properties\n", len(strs))

	strsJson, err := json.Marshal(strs)
	if err != nil {
		log.Fatalf("Error marshalling CSS properties: %v", err)
		return
	}

	os.WriteFile("properties.json", strsJson, 0644)
}

// Scrape fetches the CSS properties table from W3C and extracts the second column
// which contains the property names
func Scrape() ([]string, error) {
	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Parse HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Store the second <td> content from each row
	var properties []string

	// Find the first table and iterate through its rows
	doc.Find("table").First().Find("tr").Each(func(i int, row *goquery.Selection) {
		// Skip header row if needed
		if i > 0 || row.Find("th").Length() == 0 {
			// Get the second td in the row
			secondTd := row.Find("td").Eq(1)
			if secondTd.Length() > 0 {
				// Get text content and trim whitespace
				text := strings.TrimSpace(secondTd.Text())
				if text != "" {
					properties = append(properties, text)
				}
			}
		}
	})

	if len(properties) == 0 {
		return nil, fmt.Errorf("no properties found in the table")
	}

	return properties, nil
}

// For testing or direct usage
func scrape() {
	properties, err := Scrape()
	if err != nil {
		log.Fatalf("Error scraping CSS properties: %v", err)
	}

	fmt.Printf("Found %d CSS properties\n", len(properties))
	for i, prop := range properties {
		if i < 10 { // Print first 10 for preview
			fmt.Println(prop)
		}
	}
}
