package announcer

type Announcer interface {
	Publish(msg Message) (bool, error)
	Read() (string, error)
	Queues() ([]string, error)
}
