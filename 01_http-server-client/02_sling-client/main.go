package main

import (
	"github.com/dghubble/sling"
	"fmt"
)



func main() {
	parentSling := sling.New().Client(nil).Base("localhost:5000")
	fooSling := parentSling.New().Get("/")
	//barSling := parentSling.New().Get("bar/")

}

