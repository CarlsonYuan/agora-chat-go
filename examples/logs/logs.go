package main

import (
	"context"
	"fmt"

	agora_chat "github.com/CarlsonYuan/agora-chat-go"
	"github.com/jedib0t/go-pretty/table"
)

func main() {
	client, err := agora_chat.NewClientFromEnvVars()
	if err != nil {
		fmt.Printf("error generating token: %v\n", err)
		return
	}

	// send upload log cmd
	// client.SendUploadLogsCommand(context.Background(), "demo_user_1", "demo_user_2")

	// list log files table
	u := "demo_user_1"
	resp, _ := client.ListDeviceLogs(context.Background(), u)

	t := table.NewWriter()
	tTemp := table.Table{}
	tTemp.Render()
	fmt.Println(t.Render())
	t.AppendHeader(table.Row{"User ID", "Upload Time", "File UUID", "System Version", "SDK Version"})
	for _, item := range resp.LogFiles {

		t.AppendRow(table.Row{item.LoginUsername, item.UploadDate, item.LogfileUUID, item.OsVersion, item.SdkVersion})

		err = client.DownloadFile(item.LogfileUUID, fmt.Sprintf("tmp/%s/", item.AppKey), item.LogfileUUID+".gzip")
		if err != nil {
			fmt.Printf("err download file : %v\n", err)
		}
	}
	fmt.Println(t.Render())

}
