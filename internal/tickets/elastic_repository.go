package tickets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const ticketIndex = "tickets"

// elasticTicketRepo handles all Elasticsearch operations for tickets.
type elasticTicketRepo struct {
	es *elasticsearch.TypedClient
}

// NewElasticRepository creates a standalone Elasticsearch repository.
func NewElasticRepository(es *elasticsearch.TypedClient) ElasticRepository {
	return &elasticTicketRepo{es: es}
}

// Index indexes (or re-indexes) a ticket document using reference_no as the
// document ID.
func (r *elasticTicketRepo) Index(ctx context.Context, ticket Ticket) error {
	_, err := r.es.Index(ticketIndex).
		Id(ticket.ReferenceNo).
		Document(ticket).
		Refresh(refresh.True).
		Do(ctx)
	return err
}

// Delete removes a ticket document by its reference_no.
func (r *elasticTicketRepo) Delete(ctx context.Context, referenceNo string) error {
	_, err := r.es.Delete(ticketIndex, referenceNo).
		Refresh(refresh.True).
		Do(ctx)
	return err
}

// Search performs a full-text search across ticket fields.
func (r *elasticTicketRepo) Search(ctx context.Context, query string) ([]Ticket, error) {
	res, err := r.es.Search().
		Index(ticketIndex).
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query:  query,
				Fields: []string{"reference_no", "status"},
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch search failed: %w", err)
	}

	var tickets []Ticket
	for _, hit := range res.Hits.Hits {
		var t Ticket
		if hit.Source_ != nil {
			if err := json.Unmarshal(hit.Source_, &t); err != nil {
				return nil, fmt.Errorf("failed to unmarshal hit: %w", err)
			}
			tickets = append(tickets, t)
		}
	}

	return tickets, nil
}
