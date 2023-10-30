# agora-chat-go

Unofficial Go API client for Agora Chat.  

You can use this library to access chat API endpoints server-side.

## Installation
```
go get -u github.com/CarlsonYuan/agora-chat-go
```

## Getting started
```go
package main

import (
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
)

func main() {
    // Initialize client
    client, err := agora_chat.New("YOUR_APP_ID_HERE", "YOUR_APP_CERTIFICATE_HERE", "YOUR_BASE_URL_HERE")
    
    // Or using only environmental variables: AGORA_APP_ID, AGORA_APP_CERTIFICATE, AGORA_CHAT_BASEURL
	client, err := agora_chat.NewClientFromEnvVars()
    
    // handle error

    // Create users
	u1 := agora_chat.NewUser("self_test_1")
	u2 := agora_chat.NewUser("self_test_2")
	resp, err = client.CreateUsers(context.Background(), &u1, &u2)

    // Query a user by username(user ID)
	resp, err := client.QueryUser(context.Background(), "self_test_1")

}
```

## Contributing
You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you see any problem.