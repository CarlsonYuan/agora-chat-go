package main

import (
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {

	_, err := agora_chat.NewClientFromEnvVars()
	if err != nil {
		fmt.Printf("error generating token: %v\n", err)
		return
	}

}
