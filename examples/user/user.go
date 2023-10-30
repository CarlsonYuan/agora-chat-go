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

	// Delete User

	uID := "self_test_3"
	_, err = client.DeleteUser(context.Background(), uID)
	if err != nil {
		fmt.Printf("error deleting user: %v\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("user: %s has been deleted\n", uID)

	// Create users
	u1 := agora_chat.NewUser("self_test_3")
	u2 := agora_chat.NewUser("self_test_4")
	resp, err = client.CreateUsers(context.Background(), &u1, &u2)
	if err != nil {
		fmt.Printf("error creating users: %v\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("All list for user creation...\n")
	for _, item := range resp.Users {
		fmt.Printf(" > Create user success for %s UUID: %s\n", item.ID, item.Uuid)
	}
	for _, item := range resp.Data {
		fmt.Printf(" > Create user:%s failure cause %s \n", item.ID, item.Reason)
	}

}
