package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexandergmiles/publisher/internal/aws"
	"github.com/joho/godotenv"
	"gitlab.com/alexandergmiles/publisher/internal/announcer"
)

// Pass in JSON file containing the types of message
// and mapping to message type (version flag)
// Maybe a path to some auth details

func main() {
	fmt.Println("Starting publisher")

	messageConfig := flag.String("config", "", "Path to .JSON configuration file which maps message versions")
	envFile := flag.String("env", "", "Path to .env file")
	flag.Parse()
	print(messageConfig)

	keys, err := loadEnvFile(*envFile)
	if err != nil {
		panic(err)
	}

	for key := range keys {
		fmt.Println(keys[key])
	}
	awsConfig, err := aws.NewAWSAccess(keys["AWS_ACCESS_KEY"], keys["AWS_SECRET_ACCESS_KEY"])
	if err != nil {
		panic(err)
	}

	sqsAnnouncer := announcer.NewSQSAnnouncer(awsConfig)
	allQueues, err := sqsAnnouncer.Queues()
	if err != nil {
		panic(err)
	}

	_, err = sqsAnnouncer.Publish(allQueues[0], &announcer.BasicMessage{
		Body: "This is a message",
	})
	if err != nil {
		panic(err)
	}
}

func loadEnvFile(envFile string) (map[string]string, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, err
	}

	results := make(map[string]string, 5)

	aws_access_key := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")

	results["AWS_ACCESS_KEY"] = aws_access_key
	results["AWS_SECRET_KEY"] = aws_secret_key

	return results, nil
}
