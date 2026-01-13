package announcer

type Announcer interface {
	Subscribe() error
	Read() (string, error)
}
