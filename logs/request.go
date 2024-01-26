package logs

import (
	logging "cloud.google.com/go/logging/apiv2"
	"cloud.google.com/go/logging/apiv2/loggingpb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

type LogRequest struct {
	ProjectId string
	Filter    string
}

func ListLogEntries(logRequest LogRequest) ([]*loggingpb.LogEntry, error) {
	ctx := context.Background()

	project := fmt.Sprintf("projects/%s", logRequest.ProjectId)

	client, err := logging.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}

	defer func(client *logging.Client) {
		closErr := client.Close()
		if closErr != nil {
			log.Printf("error closing client: %w\n", closErr)
		}
	}(client)

	var entries []*loggingpb.LogEntry
	req := &loggingpb.ListLogEntriesRequest{
		ResourceNames: []string{project},
		Filter:        logRequest.Filter,
	}
	it := client.ListLogEntries(ctx, req)

	for {
		entry, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
