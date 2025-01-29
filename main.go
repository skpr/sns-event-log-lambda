package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(_ context.Context, e json.RawMessage) error {
	event, err := json.Marshal(&e)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Log this event to stdout. It will then be stored in CloudWatch Logs for future analysis.
	fmt.Print(string(event))

	return nil
}
