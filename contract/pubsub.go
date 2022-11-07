package contract

import "context"

type PubSub interface {
	SendMessage(message string) (err error)
	ReceiveMessage(ctx context.Context, chanMessage chan string) (err error)
}
