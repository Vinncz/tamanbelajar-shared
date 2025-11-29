package shared

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User domain model
type User struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey"`
	Name      string    `gorm:"not null;index"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"` // bcrypt hash
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	Wallet       *Wallet
	Transactions []Transaction
}

// Wallet domain model - one-to-one with User
type Wallet struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey"`
	UserID    uuid.UUID `gorm:"type:text;uniqueIndex;not null"`
	Balance   float64   `gorm:"not null;default:0;check:balance >= 0"`
	Version   int64     `gorm:"not null;default:0"`                   
	CreatedAt time.Time
	UpdatedAt time.Time

	User *User `gorm:"foreignKey:UserID"`
}

// Transaction types
const (
	TransactionTypeTopup       = "topup"
	TransactionTypeTransferOut = "transfer_out"
	TransactionTypeTransferIn  = "transfer_in"
	TransactionTypePayment     = "payment"
)

// Transaction statuses
const (
	TransactionStatusPending   = "pending"
	TransactionStatusCompleted = "completed"
	TransactionStatusFailed    = "failed"
	TransactionStatusExpired   = "expired"
)

// Transaction domain model
type Transaction struct {
	ID             uuid.UUID  `gorm:"type:text;primaryKey"`
	UserID         uuid.UUID  `gorm:"type:text;not null;index"`
	Type           string     `gorm:"not null;check:type IN ('topup','transfer_out','transfer_in','payment')"`
	Amount         float64    `gorm:"not null;check:amount > 0"`
	Status         string     `gorm:"not null;default:'pending';check:status IN ('pending','completed','failed','expired')"`
	FromID         *uuid.UUID `gorm:"type:text;index"` // For transfers
	ToID           *uuid.UUID `gorm:"type:text;index"` // For transfers
	From           string     `gorm:"not null"`        // Description: user UUID or system
	To             string     `gorm:"not null"`        // Description: user UUID or merchant ID
	MerchantID     *string    `gorm:"index"`           // For payments
	Description    *string							   // Optional description (such as items purchased)
	IdempotencyKey *string    `gorm:"uniqueIndex:idx_user_idempotency"` // Prevent duplicates
	ExpiresAt      *time.Time `gorm:"index"`                            // For payment expiry
	CreatedAt      time.Time  `gorm:"index"`
	UpdatedAt      time.Time
	CompletedAt    *time.Time

	// Relations
	User *User `gorm:"foreignKey:UserID"`
}

// AccessToken for tracking issued tokens
type AccessToken struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey"`
	UserID    uuid.UUID `gorm:"type:text;not null;index"`
	TokenHash string    `gorm:"not null;uniqueIndex"` // Hash of JWT for revocation
	ExpiresAt time.Time `gorm:"index"`
	IssuedAt  time.Time
	RevokedAt *time.Time

	// Relations
	User *User `gorm:"foreignKey:UserID"`
}
