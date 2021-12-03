package notification

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Selly-Modules/natsio"
)

const (
	SendByTopic = "topic"
	SendByUsers = "users"
)

const (
	SubjectPushNotification        = "push_notification"
	SubjectGetNotification         = "get_notification"
	SubjectReadNotification        = "read_notification"
	SubjectCountUnreadNotification = "count_unread_notification"
)

// Client ...
type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
}

// NewClient ...
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("api key is required")
	}
	if cfg.Nats.URL == "" {
		return nil, errors.New("nats url is required")
	}
	if err := natsio.Connect(cfg.Nats); err != nil {
		return nil, fmt.Errorf("nats connect failed: %v", err)
	}

	c := &Client{
		Config:        cfg,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
	}

	return c, nil
}

// PushToUsers push notification to list user id
func (c *Client) PushToUsers(payload PushRequest) (requestID string, err error) {
	p := pushRequest{
		APIKey:   c.Config.APIKey,
		Title:    payload.Title,
		Body:     payload.Body,
		Data:     payload.Data,
		SendBy:   SendByUsers,
		Users:    payload.Users,
		Label:    payload.Label,
		Category: payload.Category,
	}
	msg, err := c.natsServer.Request(SubjectPushNotification, toBytes(p))
	if err != nil {
		return "", err
	}
	var res struct {
		Data  PushResponse `json:"data"`
		Error string       `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &res); err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", errors.New(res.Error)
	}
	return res.Data.RequestID, nil
}

// Query get list notification by user id
func (c *Client) Query(q Query) (ListNotificationResponse, error) {
	p := query{
		APIKey:   c.Config.APIKey,
		User:     q.User,
		Category: q.Category,
		Page:     q.Page,
		Limit:    q.Limit,
	}
	msg, err := c.natsServer.Request(SubjectGetNotification, toBytes(p))
	if err != nil {
		return ListNotificationResponse{}, err
	}
	var res struct {
		Data  ListNotificationResponse `json:"data"`
		Error string                   `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &res); err != nil {
		return ListNotificationResponse{}, err
	}
	if res.Error != "" {
		return ListNotificationResponse{}, errors.New(res.Error)
	}
	return res.Data, nil
}

// CountUnread count total unread notification
func (c *Client) CountUnread(q CountUnread) (int64, error) {
	p := countUnread{
		APIKey:   c.Config.APIKey,
		User:     q.User,
		Category: q.Category,
	}
	msg, err := c.natsServer.Request(SubjectCountUnreadNotification, toBytes(p))
	if err != nil {
		return 0, err
	}
	var res struct {
		Data  CountUnreadResponse `json:"data"`
		Error string              `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &res); err != nil {
		return 0, err
	}
	if res.Error != "" {
		return 0, errors.New(res.Error)
	}
	return res.Data.Total, nil
}

// Read mark notification as read
func (c *Client) Read(notificationID string) error {
	p := read{
		APIKey: c.Config.APIKey,
		ID:     notificationID,
	}
	msg, err := c.natsServer.Request(SubjectReadNotification, toBytes(p))
	if err != nil {
		return err
	}
	var res ReadResponse
	if err := json.Unmarshal(msg.Data, &res); err != nil {
		return err
	}
	if res.Error != "" {
		err = errors.New(res.Error)
	}
	return err
}

func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
