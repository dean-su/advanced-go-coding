package subscriber

import (
	"github.com/micro/go-log"

	example "github.com/advanced-go-coding/11_micro/foo/proto/example"
	"golang.org/x/net/context"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
