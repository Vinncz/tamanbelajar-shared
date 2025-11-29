package shared

import "time"

// Event Types
const (
	EventUserCreated      = "user.created"
	EventUserUpdated      = "user.updated"
	EventTransferInitiated = "transfer.initiated"
	EventTransferCompleted = "transfer.completed"
	EventTransferFailed    = "transfer.failed"
)

// UserCreatedEvent is published when a new user registers
type UserCreatedEvent struct {
	EventID   string    `json:"eventId"`
	EventType string    `json:"eventType"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"userId"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
}

// TransferInitiatedEvent is published when transfer is requested
type TransferInitiatedEvent struct {
	EventID       string    `json:"eventId"`
	EventType     string    `json:"eventType"`
	Timestamp     time.Time `json:"timestamp"`
	TransactionID string    `json:"transactionId"`
	FromUserID    string    `json:"fromUserId"`
	ToUserID      string    `json:"toUserId"`
	Amount        float64   `json:"amount"`
}

// TransferCompletedEvent is published when transfer succeeds
type TransferCompletedEvent struct {
	EventID       string    `json:"eventId"`
	EventType     string    `json:"eventType"`
	Timestamp     time.Time `json:"timestamp"`
	TransactionID string    `json:"transactionId"`
	FromUserID    string    `json:"fromUserId"`
	ToUserID      string    `json:"toUserId"`
	Amount        float64   `json:"amount"`
}

// TransferFailedEvent is published when transfer fails
type TransferFailedEvent struct {
	EventID       string    `json:"eventId"`
	EventType     string    `json:"eventType"`
	Timestamp     time.Time `json:"timestamp"`
	TransactionID string    `json:"transactionId"`
	FromUserID    string    `json:"fromUserId"`
	ToUserID      string    `json:"toUserId"`
	Amount        float64   `json:"amount"`
	Reason        string    `json:"reason"`
}
