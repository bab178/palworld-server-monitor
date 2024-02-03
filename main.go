package main

import (
	"fmt"
	"time"
)

func main() {
	creds := getAdminData()
	output := runShowPlayers(creds)

	now := time.Now()
	playerRecords := parsePlayersOutput(output)

	fmt.Println(now.UTC().Format("2006-01-02 15:04:05"), "\n", "Players:", len(playerRecords), "\n", playerRecords)

	err := uploadRecordsToGoogleSheetsAPI(now, playerRecords)
	if err != nil {
		panic(err)
	}
}
