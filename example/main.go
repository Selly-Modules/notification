package main

import (
	"fmt"

	"github.com/Selly-Modules/natsio"

	"github.com/Selly-Modules/notification"
)

func main() {
	c, err := notification.NewClient(notification.Config{
		APIKey: "UWvieSs2erFfhwvl1g8CavEY2V7ouc3",
		Nats: natsio.Config{
			URL:      "127.0.0.1:4222",
			User:     "",
			Password: "",
			TLS:      nil,
		},
	})
	if err != nil {
		panic(err)
	}

	userID := "61a499ad8d5770f8872b03d8"
	requestID, err := c.PushToUsers(notification.PushRequest{
		Title:    "Notification 1",
		Body:     "nats stream view notification",
		Data:     "{}",
		Users:    []string{userID},
		Label:    "tracking-label",
		Category: "order",
	})
	if err != nil {
		fmt.Println("Push err: ", err)
	}
	fmt.Println("Request id: ", requestID)

	total, err := c.CountUnread(notification.CountUnread{
		User:     userID,
		Category: "order",
	})
	fmt.Println("Count: ", total, err)

	res, err := c.Query(notification.Query{
		User:     userID,
		Category: "order",
		Page:     0,
		Limit:    20,
	})
	fmt.Println("Query : ", res, err)

	total, err = c.CountUnread(notification.CountUnread{
		User:     userID,
		Category: "order",
	})
	fmt.Println("Count: 2", total, err)
}
