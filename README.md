# agora-chat-go

Unofficial Go API client for Agora Chat.  

You can use this library to access chat API endpoints server-side.

## Installation
```
go get -u github.com/CarlsonYuan/agora-chat-go
```

## Examples
## Querying a user
```go
package main

import (
	"fmt"
	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {
	// Initialize client
	client, err := agora_chat.New("YOUR_APP_ID_HERE", "YOUR_APP_CERTIFICATE_HERE", "YOUR_BASE_URL_HERE")
	if err != nil {
		fmt.Printf("error new client: %v\n", err)
		return
	}
	// Query a user by chatUid
	resp, err := client.QueryUser(context.Background(), "demo_user_1")
	if err != nil {
	fmt.Printf("error querying user: %v\n", err)
	return
	}
	user := resp.Users[0]
	fmt.Printf("\n")
	fmt.Printf("ID: %s, UUID: %s, nickname(push): %s\n", user.ID, user.Uuid, user.Nickname)
}
```
## Creating a chatroom
```go
package main

import (
	"fmt"
	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {
	// Initialize client
	client, err := agora_chat.New("YOUR_APP_ID_HERE", "YOUR_APP_CERTIFICATE_HERE", "YOUR_BASE_URL_HERE")
	if err != nil {
		fmt.Printf("error new client: %v\n", err)
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
	fmt.Printf("Chat room (ID: %s) is created \n", resp.Data.ID)
}
```

## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you see any problem.
