package imageupload

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2/event"
)


func createTestEvent(payload GCloudSEvent) (cloudevents.Event, error) {
	e := cloudevents.New()
	e.SetID("test-id")
	e.SetSource("unit-test")
	e.SetType("google.cloud.storage.object.finalize")

	data, err := json.Marshal(payload)
	if err != nil {
		return e, fmt.Errorf("failed to marshal payload: %v", err)
	}

	if err := e.SetData("application/json", data); err != nil {
		return e, fmt.Errorf("failed to set event data: %v", err)
	}
	return e, nil
}

func TestImageUploaded_WithTimestamp(t *testing.T) {
	fixedTime := "2025-03-14T12:00:00Z"
	payload := GCloudSEvent{
		Bucket: "test-bucket",
		Name: "test-image.jpg",
		TimeCreated: fixedTime,
	}
	e, err := createTestEvent(payload)
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	err = ImageUploaded(context.Background(), e)
	if err != nil {
		t.Errorf("ImageUploaded returned error: %v", err)
	}
}


func TestImageUploaded_InvalidData(t *testing.T) {
	e := cloudevents.New()
	e.SetID("invalid-id")
	e.SetSource("unit-test")
	e.SetType("google.cloud.storage.object.finalize")
	invalidData := []byte(`{"invalid": "data"}`)
	if err := e.SetData("application/json", invalidData); err != nil {
		t.Fatalf("Failed to set event data: %v", err)
	}

	err := ImageUploaded(context.Background(), e)
	if err == nil {
		t.Errorf("Expected error for invalid event data, got nil")
	}
}
