package tickets

import (
	"context"
	"database/sql"
	"gin-test/internal/tickets/db"
)

type ticketRepo struct {
	database *sql.DB
	queries  *db.Queries
}

func NewRepository(database *sql.DB) Repository {
	return &ticketRepo{
		database: database,

		// THIS IS THE INIT!
		// You pass the raw MySQL connection into sqlc's generated New() function
		queries: db.New(database),
	}
}

func (r *ticketRepo) NewTicket(ctx context.Context, referenceNo string, status string) (Ticket, error) {
	arg := db.InsertTIcketParams{
		ReferenceNo: referenceNo,
		Status:      status,
		UserID:      nil,
	}

	err := r.queries.InsertTIcket(ctx, arg)
	if err != nil {
		return Ticket{}, err
	}

	return Ticket{
		ReferenceNo: referenceNo,
		Status:      status,
	}, nil
}

func (r *ticketRepo) GetAll(ctx context.Context) ([]Ticket, error) {
	rows, err := r.queries.GetAllTickets(ctx)
	if err != nil {
		return nil, err
	}

	var tickets []Ticket
	for _, row := range rows {

		tickets = append(tickets, Ticket(row))
	}

	return tickets, nil
}

func (r *ticketRepo) Update(ctx context.Context, referenceNo string, ticket UpdateTicketRequest) (Ticket, error) {
	arg := db.UpdateTicketParams{
		ReferenceNo: referenceNo,
		Status:      ticket.Status,
	}

	err := r.queries.UpdateTicket(ctx, arg)
	if err != nil {
		return Ticket{}, err
	}

	updatedTicket, err := r.queries.GetTicketByRef(ctx, referenceNo)
	if err != nil {
		return Ticket{}, err
	}

	return Ticket(updatedTicket), nil
}
func (r *ticketRepo) Delete(ctx context.Context, referenceNo string) error {
	// 1. Capture the result now!
	result, err := r.queries.DeleteTicket(ctx, referenceNo)
	if err != nil {
		return err // This catches syntax errors or connection drops
	}

	// 2. Check how many rows MySQL actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// 3. If zero rows were deleted, the ticket didn't exist!
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
