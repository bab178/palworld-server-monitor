package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const sheetDataCSV = "sheet-data.csv"

func uploadRecordsToGoogleSheetsAPI(now time.Time, playerRecords []PlayerRecord) error {
	// Load the service account key JSON file.
	b, err := os.ReadFile("service_account.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Access Google Sheets API.
	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := conf.Client(context.Background())

	// Create a new Sheets service.
	srv, err := sheets.New(client)
	if err != nil {
		return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	// Define the data to be written.
	var vr sheets.ValueRange
	for _, record := range playerRecords {
		vr.Values = append(vr.Values, []interface{}{
			now.UTC().Format("2006-01-02 15:04:05"),
			record.Name,
			record.PlayerUID,
			record.SteamID,
		})
	}

	b, err = os.ReadFile(sheetDataCSV)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// The ID of the spreadsheet to update.
	spreadsheetId := string(b)

	// The A1 notation of the values to update.
	range2 := "A2:D2"

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, range2, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve data from sheet. %v", err)
	}

	return nil
}
