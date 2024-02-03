package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const exePath = "C:\\Servers\\ARRCON.exe"
const credCSVPath = "adminpass.csv"

type AdminDataRecord struct {
	IP       string
	Port     int
	Password string
}

func getAdminData() AdminDataRecord {
	// read ip, port, password from csv
	file, err := os.Open(credCSVPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // see the notes below
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var records []AdminDataRecord
	for _, each := range rawCSVdata {
		port, err := strconv.Atoi(each[1])
		if err != nil {
			panic(err)
		}
		record := AdminDataRecord{
			IP:       each[0],
			Port:     port,
			Password: each[2],
		}
		records = append(records, record)
	}

	return records[0]
}

func runServerCommand(creds AdminDataRecord, command string) []string {
	initialParams := []string{"-H", creds.IP, "-P", strconv.Itoa(creds.Port), "-p", creds.Password}

	params := append(initialParams, command)
	cmd := exec.Command(exePath, params...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("command failed with %s\n%s", err, output))
	}

	return strings.Split(string(output), "\r\n")
}

func runShowPlayers(creds AdminDataRecord) []string {
	return runServerCommand(creds, "ShowPlayers")
}
