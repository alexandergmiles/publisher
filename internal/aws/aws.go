package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AWSAccess struct{}

func NewAWSAccess(accessKey string, secretKey string) (*aws.Config, error) {
	timeout, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	config, err := config.LoadDefaultConfig(timeout)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate with AWS: %s", err)
	}
	return &config, nil
}
