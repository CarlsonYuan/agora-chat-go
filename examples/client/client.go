package main

import (
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {

	_, err := agora_chat.New("YOUR_APP_ID_HERE", "YOUR_APP_CERTIFICATE_HERE", "YOUR_BASE_URL_HERE")
	if err != nil {
		fmt.Printf("error generating token: %v\n", err)
		return
	}

}
