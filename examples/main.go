package main

import (
	"flag"
	"fmt"

	"github.com/Lexographics/notifywork-go"
)

var apiKeyFlag = flag.String("key", "", "api key")
var messageFlag = flag.String("m", "Test Message", "message to send")
var channelFlag = flag.String("c", "", "channel id")
var senderId = flag.Uint("sender-id", 0, "ID of the sender")

func main() {
	flag.Parse()

	sender := notifywork.NewSender(*apiKeyFlag, *senderId)
	sender.SetDefaultChannel(*channelFlag)
	err := sender.SendMessage(*messageFlag)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("Message sent")
	}
}
