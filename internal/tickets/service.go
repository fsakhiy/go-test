package tickets

import (
	"context"
	"fmt"
	// "errors"
)

type ticketService struct {
	repo    Repository
	elastic ElasticRepository
}

func NewService(repo Repository, elastic ElasticRepository) Service {
	return &ticketService{repo: repo, elastic: elastic}
}

func (s *ticketService) Test(ctx context.Context) (string, error) {
	return "Hello, Test!", nil
}

func (s *ticketService) NewTicket(ctx context.Context, ticket Ticket) (Ticket, error) {
	// 1. Persist in MySQL (source of truth)
	created, err := s.repo.NewTicket(ctx, ticket.ReferenceNo, ticket.Status, *ticket.UserID)
	if err != nil {
		return Ticket{}, err
	}

	// 2. Index to Elasticsearch
	if err := s.elastic.Index(ctx, created); err != nil {
		fmt.Printf("Warning: failed to index ticket %s in Elasticsearch: %v\n", ticket.ReferenceNo, err)
	}

	return created, nil
}

func (s *ticketService) GetAll(ctx context.Context) ([]Ticket, error) {
	return s.repo.GetAll(ctx)
}

func (s *ticketService) Update(ctx context.Context, referenceNo string, ticket UpdateTicketRequest) (Ticket, error) {
	// 1. Update in MySQL
	updated, err := s.repo.Update(ctx, referenceNo, ticket)
	if err != nil {
		return Ticket{}, err
	}

	// 2. Re-index in Elasticsearch
	if err := s.elastic.Index(ctx, updated); err != nil {
		fmt.Printf("Warning: failed to re-index ticket %s in Elasticsearch: %v\n", referenceNo, err)
	}

	return updated, nil
}

func (s *ticketService) Delete(ctx context.Context, referenceNo string) error {
	// 1. Delete from MySQL
	if err := s.repo.Delete(ctx, referenceNo); err != nil {
		return err
	}

	// 2. Remove from Elasticsearch
	if err := s.elastic.Delete(ctx, referenceNo); err != nil {
		fmt.Printf("Warning: failed to delete ticket %s from Elasticsearch: %v\n", referenceNo, err)
	}

	return nil
}
