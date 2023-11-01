package main

import (
	"context"
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {

	client, err := agora_chat.NewClientFromEnvVars()
	if err != nil {
		fmt.Printf("error create chat client: %v\n", err)
		return
	}

	// APNS P8
	/*
		p := agora_chat.PushProvider{
			Type:             agora_chat.PushProviderAPNS,
			Name:             "test2-APPLE-DEVELOPMENT-p8",
			Env:              agora_chat.EnvDevelopment,
			Certificate:      "XXX",
			PackageName:      "com.hyphenate.easeim",
			ApnsPushSettings: &agora_chat.APNSConfig{TeamId: "499WYUV8Q2", KeyId: "B6FZWPWS4L"},
		}
		client.InsertPushProvider(context.Background(), &p)
	*/

	// List of push providers
	resp, err := client.ListPushProviders(context.Background())
	if err != nil {
		fmt.Printf("error get providers list: %v\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("All push provider list \n")
	for _, item := range resp.PushProviders {
		fmt.Printf("Item type: %s\n", item.Type)
		fmt.Printf(" > Item uuid: %s\n", item.UUID)
		fmt.Printf(" > Item name: %s\n", item.Name)
		fmt.Printf(" > Item packageName: %s\n", item.PackageName)
	}

	// Delete a push provider
	_, err = client.DeletePushProvider(context.Background(), resp.PushProviders[0].UUID)
	if err != nil {
		fmt.Printf("error delete provider: %v\n", err)
		return
	}
}
