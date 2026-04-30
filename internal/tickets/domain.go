package tickets

import (
	"context"
	"time"
)

// 1. The Struct
type Ticket struct {
	ID          int64  `json:"id"`
	ReferenceNo string `json:"reference_no" validate:"required"`
	Status      string `json:"status" validate:"required"`
	// Use a pointer for nullable integers
	UserID *int64 `json:"user_id"`
	// Use a pointer for nullable times
	CreatedAt *time.Time `json:"created_at"`
}

// 2. The Repository Interface (What the DB must do)
type Repository interface {
	// ClaimTicket(ctx context.Context, ticketID int64, userID int64) error
	NewTicket(ctx context.Context, referenceNo string, status string) (Ticket, error)
	GetAll(ctx context.Context) ([]Ticket, error)
	Update(ctx context.Context, referenceNo string, ticket UpdateTicketRequest) (Ticket, error)
	Delete(ctx context.Context, referenceNo string) error
}

// 3. The Service Interface (What the business logic must do)
type Service interface {
	// BuyTicket(ctx context.Context, ticketID int64, userID int64) error
	Test(ctx context.Context) (string, error)

	NewTicket(ctx context.Context, ticket Ticket) (Ticket, error)
	GetAll(ctx context.Context) ([]Ticket, error)
	Update(ctx context.Context, referenceNo string, ticket UpdateTicketRequest) (Ticket, error)
	Delete(ctx context.Context, referenceNo string) error
}
