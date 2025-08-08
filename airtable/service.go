package airtable

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const baseURL = "https://api.airtable.com/v0/"

func getToken() string {
	t := os.Getenv("AIRTABLE_TOKEN")
	if t == "" {
		panic("AIRTABLE_TOKEN not set")
	}
	return "Bearer " + t
}

type AirtableService struct{}

func (s *AirtableService) GetAirtableData(w http.ResponseWriter, r *http.Request, baseId string, tableName string) {
	url := fmt.Sprintf("%s%s/%s", baseURL, baseId, tableName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", getToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Request Failed", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

func (s *AirtableService) GetAirtableRecord(w http.ResponseWriter, r *http.Request, baseId string, tableName string, recordId string) {
	url := fmt.Sprintf("%s%s/%s/%s", baseURL, baseId, tableName, recordId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", getToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Request Failed", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

func (s *AirtableService) GetRecord(ctx echo.Context, baseID string, tableName string, recordID string) error {
	url := fmt.Sprintf("%s%s/%s/%s", baseURL, baseID, tableName, recordID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to build request"})
	}
	req.Header.Set("Authorization", getToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Request to Airtable failed"})
	}
	defer res.Body.Close()

	// Read Airtable response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read Airtable response"})
	}

	// Return the response with the original status code from Airtable
	return ctx.Blob(res.StatusCode, "application/json", body)
}

func (s *AirtableService) ListRecords(ctx echo.Context, baseID string, tableName string) error {
	url := fmt.Sprintf("%s%s/%s", baseURL, baseID, tableName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
	}
	req.Header.Set("Authorization", getToken())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Request Failed"})
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response"})
	}

	// Set the response headers and return raw Airtable response
	return ctx.Blob(res.StatusCode, res.Header.Get("Content-Type"), body)
}
