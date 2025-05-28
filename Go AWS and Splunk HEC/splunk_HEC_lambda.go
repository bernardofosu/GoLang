package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	SPLUNK_HEC   = "https://<splunk-host>:8088/services/collector"
	SPLUNK_TOKEN = "<your-hec-token>"
)

// CloudTrailEvent represents the structure of one record in CloudTrail logs
type CloudTrailEvent struct {
	// Use map[string]interface{} if you want a generic dynamic structure
	// or define fields if you want typed access
	// Here we keep it generic
	// e.g. EventVersion string `json:"eventVersion"`
	// For simplicity:
	Raw map[string]interface{} `json:"-"`
}

func handler(ctx context.Context, s3Event events.S3Event) error {
	// Create AWS session
	sess := session.Must(session.NewSession())
	s3Client := s3.New(sess)

	httpClient := &http.Client{}

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		// Get the S3 object
		objOutput, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			log.Printf("Failed to get object from S3: %v", err)
			return err
		}
		defer objOutput.Body.Close()

		// Decompress gzip
		gzipReader, err := gzip.NewReader(objOutput.Body)
		if err != nil {
			log.Printf("Failed to create gzip reader: %v", err)
			return err
		}
		defer gzipReader.Close()

		// Read decompressed JSON bytes
		decompressedBytes, err := io.ReadAll(gzipReader)
		if err != nil {
			log.Printf("Failed to read decompressed data: %v", err)
			return err
		}

		// Parse JSON into a generic map
		var cloudTrailData map[string][]map[string]interface{}
		err = json.Unmarshal(decompressedBytes, &cloudTrailData)
		if err != nil {
			log.Printf("Failed to unmarshal JSON: %v", err)
			return err
		}

		records, ok := cloudTrailData["Records"]
		if !ok {
			log.Printf("No Records found in JSON")
			continue
		}

		// Send each event to Splunk HEC
		for _, eventEntry := range records {
			// Prepare payload
			payloadMap := map[string]interface{}{
				"event":      eventEntry,
				"sourcetype": "aws:cloudtrail",
				"index":      "cloudtrail",
				"host":       "aws",
			}
			payloadBytes, err := json.Marshal(payloadMap)
			if err != nil {
				log.Printf("Failed to marshal payload: %v", err)
				continue
			}

			// Create HTTP request
			req, err := http.NewRequest("POST", SPLUNK_HEC, bytes.NewBuffer(payloadBytes))
			if err != nil {
				log.Printf("Failed to create HTTP request: %v", err)
				continue
			}
			req.Header.Set("Authorization", "Splunk "+SPLUNK_TOKEN)
			req.Header.Set("Content-Type", "application/json")

			// Send request
			resp, err := httpClient.Do(req)
			if err != nil {
				log.Printf("Failed to send request to Splunk HEC: %v", err)
				continue
			}

			log.Printf("HEC status: %s", resp.Status)
			resp.Body.Close()
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
