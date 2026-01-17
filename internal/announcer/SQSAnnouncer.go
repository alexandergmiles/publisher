package announcer

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSAnnouncer struct {
	SQSClient sqs.Client
}

func NewSQSAnnouncer(config *aws.Config) *SQSAnnouncer {
	announcer := SQSAnnouncer{}

	sqsClient := sqs.NewFromConfig(*config)
	announcer.SQSClient = *sqsClient

	return &announcer
}

func (s *SQSAnnouncer) Publish(queue string, msg Message) (bool, error) {
	content, err := msg.Content()
	if err != nil {
		return false, fmt.Errorf("unable to get message body: %s", err)
	}
	_, err = s.SQSClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &queue,
		MessageBody: &content,
	})
	if err != nil {
		return false, fmt.Errorf("unable to send message: %s", err)
	}

	return true, nil
}

func (s *SQSAnnouncer) Queues() ([]string, error) {
	timeout, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	allQueueURLs := []string{}

	pag := sqs.NewListQueuesPaginator(&s.SQSClient, nil)

	for pag.HasMorePages() {
		currentPage, err := pag.NextPage(timeout)
		if err != nil {
			return allQueueURLs, fmt.Errorf("unable to get all queue URLs: %s", err)
		}
		allQueueURLs = append(allQueueURLs, currentPage.QueueUrls...)
	}

	return allQueueURLs, nil
}
