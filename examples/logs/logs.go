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

	// send upload log cmd
	client.SendUploadLogsCommand(context.Background(), "self_test_1", "self_test_2")

	// list log files
	client.ListDeviceLogs(context.Background(), "self_test_1")
}
