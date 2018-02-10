package main

import (
	"github.com/parnurzeal/gorequest"
	//"bytes"
	//"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     string `json:"id"`
}

func main() {
	//var u User

	request := gorequest.New()
	resp, body, errs := request.Post("http://localhost:5000").
		Set("Notes","gorequst is coming!").
		Send(`{"Username":"testuser", "Password":"testpass"}`).
		End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode)
	fmt.Printf("\nResponse Status: %v", resp.Status)
	fmt.Printf("\nResponse Header: %v", resp.Header)
	fmt.Printf("\nResponse ContentLength : %v", resp.ContentLength)
	fmt.Printf("\nResponse Body: %v", resp.Body)
/*	err = json.Unmarshal(resp.Body, &u)
	if err != nil {
		fmt.Println("error:", err)
	}*/
	fmt.Printf("%+v", body)
}