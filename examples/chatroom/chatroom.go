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

	// create a chat room
	/*
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
	*/

	// retrieves the basic information of all chat rooms under the app by page.
	//*
	req := agora_chat.QueryChatroomsRequest{
		Pagination: &agora_chat.PaginationParamsRequest{
			Limit:  10,
			Cursor: "",
		},
	}
	resp, err := client.QueryChatrooms(context.Background(), &req)
	if err != nil {
		fmt.Printf("error retrieving chatrooms: %v\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("All chat rooms under the app\n")
	for _, item := range resp.Data {
		fmt.Printf(" > ID: %s, Name: %s\n", item.ID, item.Name)
	}
	/*/

	// retrieves all the chat rooms that a user joins.
	/*
		req := agora_chat.QueryChatroomsRequest{
			Pagination: &agora_chat.PaginationParamsRequest{
				Limit:  10,
				Cursor: "",
			},
		}
		resp, err := client.QueryUserJoinedChatrooms(context.Background(), "demo_user_1", &req)
		if err != nil {
			fmt.Printf("error retrieving chatroom: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("All chat rooms under that a user joins\n")
		for _, item := range resp.Data {
			fmt.Printf(" > ID: %s, Name: %s\n", item.ID, item.Name)
		}
	*/

	// retrieves the detailed information of one or more specified chat rooms.
	/*
		roomID1 := "231256978685954"
		roomID2 := "231257325764611"
		resp, err := client.QuerySpecifiedChatrooms(context.Background(), roomID1, roomID2)
		if err != nil {
			fmt.Printf("error retrieving chatroom: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Detailed information of Chat rooms\n")
		for _, item := range resp.Data {
			fmt.Printf(" > ID: %s, Name: %s\n", item.ID, item.Name)
		}
	*/

	// roomID := "231256978685954" // for testing

	// Modifies the information of the specified chat room.
	// You can only modify the name, description, and maxusers of a chat room
	/*
		req := agora_chat.UpdateChatroomRequest{
			Name:        "demo chatroom name",
			Description: "a description",
			Maxusers:    100,
		}
		resp, err := client.UpdateChatroom(context.Background(), "231256978685954", &req)
		if err != nil {
			fmt.Printf("error retrieving chatroom: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Chat room info has been modified\n")
		for key, value := range *resp.Data {
			fmt.Printf("> %s value is %v\n", key, value)
		}
	*/

	// Deletes the specified chat room
	/*
		roomID := "231879130284035"
		_, err = client.DeleteChatroom(context.Background(), roomID)
		if err != nil {
			fmt.Printf("error retrieving chatroom: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Chat room (ID: %s) has been deleted\n", roomID)
	*/

	// Retrieves the announcement text for the specified chat room.
	/*
		resp, err := client.QueryChatroomAnnouncement(context.Background(), roomID)
		if err != nil {
			fmt.Printf("error retrieving chatroom announcement: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Chat room announcement is %v\n", resp.Data)
	*/

	// Modifies the announcement text of the specified chat room.
	/*
		req := agora_chat.UpdateChatroomAnnouncementRequest{
			Announcement: "this is a demo announcement",
		}
		resp, err := client.UpdateChatroomAnnouncement(context.Background(), roomID, &req)
		if err != nil {
			fmt.Printf("error updating chatroom announcement: %v\n", err)
			return
		}
		fmt.Printf("\n")
		fmt.Printf("Chat room announcement is updated = %v\n", resp.Data.Result)
	*/
}
