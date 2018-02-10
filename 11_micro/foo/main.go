package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/advanced-go-coding/11_micro/foo/handler"
	"github.com/advanced-go-coding/11_micro/foo/subscriber"

	example "github.com/advanced-go-coding/11_micro/foo/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.example.srv.foo"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.com.example.srv.foo", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.com.example.srv.foo", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
