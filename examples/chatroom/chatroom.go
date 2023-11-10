package main

import (
	"context"
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {

	client, err := agora_chat.NewClientFromEnvVars()

	if err != nil {
		fmt.Printf("error generating token: %v\n", err)
		return
	}

	// create a chat room
	cr := agora_chat.Chatroom{
		Name:        "demo chatroom name",
		Description: "a description",
		Maxusers:    100,
		Owner:       "demo_user_1",
		Members:     []string{"demo_user_1", "demo_user_2"},
	}

	resp, err := client.CreateChatroom(context.Background(), &cr)
	if err != nil {
		fmt.Printf("error creating chatroom: %v\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("%v\n", resp.Data)
}
