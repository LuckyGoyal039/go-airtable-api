package main

import (
	"fmt"
	"io"
	"net/http"
)

const airtableBaseURL = "https://api.airtable.com/v0/"
const airtableToken = "pat1EaMfMyRi6scoE.97d36d16f22098570d408c2fe24ab2d3cd4205ab36db4ab427008698bf22a374"

func main() {
	// Simulate input
	baseID := "appOsM0fcKAqWmxca"
	tableName := "Contacts"
	recordID := "recp6PGKakewf9IpW" // leave empty for list, or set to an ID like "rec123..."

	if recordID == "" {
		listRecords(baseID, tableName)
	} else {
		getRecord(baseID, tableName, recordID)
	}
}

func listRecords(baseID, tableName string) {
	url := fmt.Sprintf("%s%s/%s", airtableBaseURL, baseID, tableName)
	fmt.Println("GET:", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+airtableToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:")
	fmt.Println(string(body))
}

func getRecord(baseID, tableName, recordID string) {
	url := fmt.Sprintf("%s%s/%s/%s", airtableBaseURL, baseID, tableName, recordID)
	fmt.Println("GET:", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Request creation failed:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+airtableToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:")
	fmt.Println(string(body))
}
