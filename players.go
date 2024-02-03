package main

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerRecord struct {
	Name      string
	PlayerUID int64
	SteamID   int64
}

func parsePlayersOutput(lines []string) []PlayerRecord {
	var playerRecords []PlayerRecord

	foundColumnHeadersLine := false
	for _, line := range lines {
		// skip all lines until we find the column headers
		if !foundColumnHeadersLine {
			if strings.Contains(line, "name,playeruid,steamid") {
				foundColumnHeadersLine = true
			}
			continue
		}

		// skip lines that don't contain a comma
		if !strings.Contains(line, ",") {
			continue
		}

		// these lines are just player data
		playerRecords = append(playerRecords, parsePlayerRecord(line))
	}
	return playerRecords
}

func parsePlayerRecord(line string) PlayerRecord {
	parts := strings.Split(line, ",")

	if len(parts) != 3 {
		panic(fmt.Sprintf("expected 3 parts, got %d, %s", len(parts), line))
	}

	playerUID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}

	steamID, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		panic(err)
	}

	return PlayerRecord{
		Name:      parts[0],
		PlayerUID: playerUID,
		SteamID:   steamID,
	}
}
