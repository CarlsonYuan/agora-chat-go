## About agora-chat-go
agora-chat-go is a Go Client for [AgoraChat](https://docs.agora.io/en/agora-chat/overview/product-overview?platform=android).  

You can refer to this library when you need to manage AgoraChat API server side.

### Creating a Client
```
import chat "github.com/CarlsonYuan/agora-chat-go/v2"

appID := "YOUR_APPID"
appCertificate := "YOUR_APPCERTIFICATE"
baseURL := "https://YOUR_REST_API/YOUR_ORG_NAME/YOUR_APP_NAME"
client, err := chat.NewClient(appID, appCertificate, baseURL)
if err != nil {
    // ...
}
```

### Generating Tokens

When using the create user token method, pass the user_ID parameter to generate a client-side token.
```
// creates a token valid for 2 hour for user "wukong"
userToken, err := client.CreateUserToken("wukong", 2*60*60)
```

### Users
Create users
```
import chat "github.com/CarlsonYuan/agora-chat-go/v2"
result, err := client.CreateUsers(ctx, &chat.User{Username: "test_user_1", Password: "1", Nickname: "test_user_1"},
    &chat.User{Username: "test_user_2", Password: "2", Nickname: "test_user_2"},
    &chat.User{Username: "test_user_3", Password: "3", Nickname: "test_user_3"})
```
Delete user
```
result, err := client.DeleteUser(ctx, "tese_user_1")
```
