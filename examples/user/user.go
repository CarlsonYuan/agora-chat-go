package main

import (
	"context"
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {

	client, err := agora_chat.New("YOUR_APP_ID_HERE", "YOUR_APP_CERTIFICATE_HERE", "YOUR_BASE_URL_HERE")

	if err != nil {
		fmt.Printf("error generating token: %v\n", err)
		return
	}

	// 1. Query a user by username(user ID)
	resp, err := client.QueryUser(context.Background(), "wukong")
	if err != nil {
		fmt.Printf("error querying user: %v\n", err)
		return
	}
	user := resp.Users[0]
	fmt.Printf("\n")
	fmt.Printf("ID: %s, UUID: %s, nickname(push): %s\n", user.ID, user.Uuid, user.Nickname)

	// 1.1 List all of the users push info.
	pushInfos := user.PushInfos
	fmt.Printf("\n")
	fmt.Printf("All push info list by %s...\n", user.ID)
	for _, item := range pushInfos {
		fmt.Printf(" > Item NotifierName: %s\n", item.NotifierName)
	}

}
