package aws

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	envFile        string
	aws_secret_key string
	aws_key_id     string
)

func TestMain(m *testing.M) {
	resultEnvFile := flag.String("env", "", "Env file for testing")
	flag.Parse()

	envFile = *resultEnvFile

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("Unable to load env file: %s\n", err)
		os.Exit(-1)
	}

	aws_secret_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	aws_key_id = os.Getenv("AWS_ACCESS_KEY_ID")

	// Kick off test run
	m.Run()
}

func TestAuthenticate(t *testing.T) {
	_, err := NewAWSAccess(aws_key_id, aws_secret_key)
	if err != nil {
		t.Errorf("unable to authenticate: %s\n", err)
	}
}
