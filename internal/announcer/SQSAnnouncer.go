package announcer

import "fmt"

type QueueTopic struct {
	Name string
}

type SQSAnnouncer struct {
	Topic QueueTopic
}

func (S *SQSAnnouncer) Subscribe() error {
	return fmt.Errorf("unimplemented")
}

func (S *SQSAnnouncer) Read() (string, error) {
	return "", fmt.Errorf("unimplemented")
}
