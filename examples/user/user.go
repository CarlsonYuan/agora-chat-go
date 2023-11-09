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

	// Create users
	/*
		u1 := agora_chat.NewUser("demo_user_1")
		u2 := agora_chat.NewUser("demo_user_2")
		resp, err := client.CreateUsers(context.Background(), &u1, &u2)
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
	*/

	// Delete User
	/*
		uID := "demo_user_2"
		_, err = client.DeleteUser(context.Background(), uID)
		if err != nil {
			fmt.Printf("error deleting user: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("user: %s has been deleted\n", uID)
	*/

	// add push info
	pushInfo := agora_chat.PushInfo{
		DeviceID:     "6cfce61c-b9b0-f4e7-c5fd-cb8db66a51ae",
		DeviceToken:  "80ceb30796ddd65c4f58286c2af17b85351bd580eb8a0c1485fbc2661d46c0f9",
		NotifierName: "AgoraChatDemoDevPush",
	}
	/* Uncomment to delete pushInfo
	pushInfo := agora_chat.PushInfo{
		DeviceID: "6cfce61c-b9b0-f4e7-c5fd-cb8db66a51ae",
		// DeviceToken:  "4569dbb268733126f4b4502beced12740a8dfe770879ea105959e4bb7de4d58d",
		DeviceToken:  "",
		NotifierName: "AgoraChatDemoDevPush",
	}
	*/
	_, err = client.UpdatePushInfo(context.Background(), "demo_user_1", &pushInfo)
	if err != nil {
		fmt.Printf("error add pushInfo: %v\n", err)
		return
	}

	// Example query a user detail
	exampleQueryUser(client)
}

// Example query a user detail
func exampleQueryUser(client *agora_chat.Client) {
	// Query a user by username(user ID)
	resp, err := client.QueryUser(context.Background(), "demo_user_1")
	if err != nil {
		fmt.Printf("error querying user: %v\n", err)
		return
	}
	user := resp.Users[0]
	fmt.Printf("\n")
	fmt.Printf("ID: %s, UUID: %s, nickname(push): %s\n", user.ID, user.Uuid, user.Nickname)

	// List all of the users push info.
	pushInfos := user.PushInfos
	fmt.Printf("\n")
	fmt.Printf("All push info list by %s...\n", user.ID)
	for _, item := range pushInfos {
		fmt.Printf(" > Item NotifierName: %s\n", item.NotifierName)
		fmt.Printf(" > Item DeviceID: %s\n", item.DeviceID)
		fmt.Printf(" > Item DeviceToken: %s\n", item.DeviceToken)
	}

}
