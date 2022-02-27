package main

import (
	"context"
	"log"

	"github.com/aflmp/ping-pong-function/lambda-function/ping"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	var config ping.Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatalf("failed to process config: %v", err)
	}

	lambda.Start(config.Ping)
}
