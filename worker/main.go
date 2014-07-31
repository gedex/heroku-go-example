package main

import (
	"fmt"
	"github.com/gedex/heroku-go-example/message"
	"time"
)

func main() {
	for {
		fmt.Println(message.Hello)
		time.Sleep(10 * time.Second)
	}
}
