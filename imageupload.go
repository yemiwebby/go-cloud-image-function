package imageupload

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2/event"
)


type GCloudSEvent struct {
	Bucket string `json:"bucket"`
	Name string `json:"name"`
	Metageneration string `json:"metageneration"`
	TimeCreated string `json:"timeCreated"`
	Updated string `json:"updated"`
}

func init() {
	// Register the CloudEvent
	functions.CloudEvent("ImageUploaded", ImageUploaded)
}


func ImageUploaded(ctx context.Context, e cloudevents.Event) error {
	var payload GCloudSEvent

	if err := e.DataAs(&payload); err != nil {
		return fmt.Errorf("failed to parse event data: %v", err)
	}


	timestamp := payload.TimeCreated
	if timestamp == "" {
		timestamp = time.Now().Format(time.RFC3339)
	}
	fmt.Printf("Image '%s' has just been uploaded to bucket '%s' at %s\n", payload.Name, payload.Bucket, timestamp)
	return nil
}