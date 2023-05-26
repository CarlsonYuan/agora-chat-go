package agora_chat

type MessageType string

const (
	MessageTypeTxt   MessageType = "txt"
	MessageTypeImg   MessageType = "img"
	MessageTypeAudio MessageType = "audio"
	MessageTypeVideo MessageType = "video"
	MessageTypeFile  MessageType = "file"
	MessageTypeLoc   MessageType = "loc"
	MessageTypeCmd   MessageType = "cmd"
)

type RouteType string

const (
	RouteTypeOnline RouteType = "ROUTE_ONLINE"
)

type Attachment struct { // type = img audio video file
	Filename    string `json:"filename"`
	Secret      string `json:"secret"`
	Url         string `json:"url"`
	Thumb       string `json:"thumb"`
	ThumbSecret string `json:"thumb_secret"`
	Length      int    `json:"length"`
	FileLength  int    `json:"file_length"`
	Size        struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"size"`
	RouteType  string `json:"routetype"`
	SyncDevice bool   `json:"sync_device"`
}

//"type": "loc","body":{"lat": "39.966","lng":"116.322","addr":"North America"}}'
//"type": "cmd","body":{"action":"action1"}}'

type Message struct {
	From        string   `json:"from"` // the user ID of the sender
	To          []string `json:"to"`   // the user IDs of the receiver
	MessageType string   `json:"type"`
	Body        struct {
		Msg string `json:"msg"`
		Attachment
	} `json:"body"`
}

// SendMessage sends a message to the channel. Returns full message details from server.
//func (usr *User) SendMessage(ctx context.Context, message *Message, userID string, options ...SendMessageOption) (*MessageResponse, error) {
//	switch {
//	case message == nil:
//		return nil, errors.New("message is nil")
//	case userID == "":
//		return nil, errors.New("user ID must be not empty")
//	}
//
//	message.User = &User{ID: userID}
//	p := path.Join("channels", url.PathEscape(ch.Type), url.PathEscape(ch.ID), "message")
//
//	req := message.toRequest()
//	for _, op := range options {
//		op(&req)
//	}
//
//	var resp MessageResponse
//	err := ch.client.makeRequest(ctx, http.MethodPost, p, nil, req, &resp)
//	return &resp, err
//}
