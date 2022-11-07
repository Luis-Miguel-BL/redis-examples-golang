package contract

type Queue interface {
	SendMessage(msg string) (err error)
	ReceiveMessage() (msg string, err error)
}
