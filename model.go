package notification

import "github.com/Selly-Modules/natsio"

// Config ...
type Config struct {
	APIKey string
	Nats   natsio.Config
}

// PushRequest ...
type PushRequest struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Data  string   `json:"data"`
	Users []string `json:"users"`
	Label string   `json:"label"`
}

// PushResponse ...
type PushResponse struct {
	RequestID string `json:"requestId"`
	Error     string `json:"error"`
}

type pushRequest struct {
	APIKey string   `json:"apiKey"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Data   string   `json:"data"`
	SendBy string   `json:"sendBy"`
	Users  []string `json:"users"`
	Topic  string   `json:"topic"`
	Label  string   `json:"label"`
}

// Query ...
type Query struct {
	User     string `json:"user"`
	Category string `json:"category,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
}

type query struct {
	APIKey   string `json:"apiKey"`
	User     string `json:"user"`
	Category string `json:"category,omitempty"`
	Page     int64  `json:"page,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
}

// Notification ...
type Notification struct {
	ID         string `json:"id"`
	Category   string `json:"category,omitempty"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	IsRead     bool   `json:"isRead"`
	Data       string `json:"data,omitempty"`
	CreatedAt  string `json:"createdAt"`
	LastPushAt string `json:"lastPushAt"`
}

// ListNotificationResponse ...
type ListNotificationResponse struct {
	List  []Notification `json:"list"`
	Total int64          `json:"total"`
	Limit int64          `json:"limit"`
}

// Read ...
type Read struct {
	APIKey string `json:"apiKey"`
	ID     string `json:"id"`
}

// ReadResponse ...
type ReadResponse struct {
	Error string `json:"error"`
}

type read struct {
	APIKey string `json:"apiKey"`
	ID     string `json:"id"`
}

// CountUnread ...
type CountUnread struct {
	User     string `json:"user"`
	Category string `json:"category"`
}

// CountUnreadResponse ...
type CountUnreadResponse struct {
	Total int64 `json:"total"`
}

type countUnread struct {
	APIKey   string `json:"apiKey"`
	User     string `json:"user"`
	Category string `json:"category"`
}
