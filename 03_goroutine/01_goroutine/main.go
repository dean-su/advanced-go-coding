package main

import (
	"fmt"
	//"runtime"
	"time"
)

func say(s string)  {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}

