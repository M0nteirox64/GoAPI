package main

import (
	"github.com/go-resty/resty/v2"
	"fmt"
)


func main() {
	client := resty.New()
	var msg string

	fmt.Print("POST: ")
	fmt.Scan(&msg)


	resp, err := client.R().
		SetBody(map[string]string{
			"message": msg,
		}).
		Post("http://localhost:8080/message")

	if err != nil {
		fmt.Println(err)
	}
	

	if resp.StatusCode() == 200 {
		fmt.Println("[200] OK")
	}
}
