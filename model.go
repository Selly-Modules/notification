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

type Read struct {
	APIKey string `json:"apiKey"`
	ID     string `json:"id"`
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

type countUnread struct {
	APIKey   string `json:"apiKey"`
	User     string `json:"user"`
	Category string `json:"category"`
}