package main

import (
	"bufio"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"encoding/json"
	"extractlogs/logs"
	"fmt"
	"log"
	"os"
)

func WriteSliceToFile(slice []*loggingpb.LogEntry, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}(file)

	writer := bufio.NewWriter(file)

	jsonData, err := json.Marshal(slice)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing to file: %w", err)
	}

	return nil
}

func main() {

	filter := logs.NewFilterBuilder()
	filter.BySeverity("ERROR").ByText("some error").CustomFilter("resource.type", "=", "gce_instance").Build()

	logRequest := logs.LogRequest{
		ProjectId: "project-id	",
		Filter:    filter.Build(),
	}

	entries, err := logs.ListLogEntries(logRequest)
	if err != nil {
		log.Fatalf("Failed to get log entries: %v", err)
	}

	err = WriteSliceToFile(entries, "logentries.txt")

	if err != nil {
		log.Fatalf("Failed to write log entries to file: %v", err)
	}

}
