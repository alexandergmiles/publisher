package consumer

type Consumer interface {
	Read() (string, error)
}
