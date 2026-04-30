package tickets

import (
	"context"
	// "errors"
)

type ticketService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &ticketService{repo: repo}
}

func (s *ticketService) Test(ctx context.Context) (string, error) {
	return "Hello, Test!", nil
}

func (s *ticketService) NewTicket(ctx context.Context, ticket Ticket) (Ticket, error) {
	return s.repo.NewTicket(ctx, ticket.ReferenceNo, ticket.Status)
}

func (s *ticketService) GetAll(ctx context.Context) ([]Ticket, error) {
	return s.repo.GetAll(ctx)
}

func (s *ticketService) Update(ctx context.Context, referenceNo string, ticket UpdateTicketRequest) (Ticket, error) {
	return s.repo.Update(ctx, referenceNo, ticket)
}

func (s *ticketService) Delete(ctx context.Context, referenceNo string) error {
	return s.repo.Delete(ctx, referenceNo)
}
