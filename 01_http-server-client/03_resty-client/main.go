package main

import (
	"gopkg.in/resty.v1"
	"fmt"
	"time"
	"encoding/json"
	"os"
	"os/signal"

	"log"
	"context"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Client is starting...")
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		//atomic.StoreInt32(&int32, 0)

		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()


		close(done)
	}()

	// GET request
	var u User

	for {
		resp, err := resty.R().SetHeader("Content-Type", "application/json").SetBody(`{"Username":"testuser", "Password":"testpass"}`).Post("http://localhost:5000")
		if err != nil {
			fmt.Println("error:", err)
		}
		// explore response object
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v", string(resp.Body()))
		err = json.Unmarshal(resp.Body(), &u)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%+v", u)
		time.Sleep(time.Millisecond * 500)
	}
}
