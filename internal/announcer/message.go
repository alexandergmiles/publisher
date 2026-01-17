package announcer

type Message interface {
	Encode(body string) error
	Content() (string, error)
}

type BasicMessage struct {
	Body    string
	Version int
}

func (b *BasicMessage) Encode(body string) error {
	return nil
}

func (b *BasicMessage) Content() (string, error) {
	return b.Body, nil
}
