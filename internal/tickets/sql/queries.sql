-- name: GetTicketByRef :one
SELECT * FROM tickets WHERE reference_no = ? LIMIT 1;

-- name: GetAllTickets :many
SELECT * FROM tickets;

-- name: ClaimTicket :exec
UPDATE tickets SET status = ?, user_id = ? WHERE id = ?;

-- name: InsertTIcket :exec
INSERT INTO tickets (reference_no, status, user_id) VALUES (?, ?, ?);

-- name: UpdateTicket :exec
UPDATE tickets SET status = ? WHERE reference_no = ?;

-- name: DeleteTicket :execresult
DELETE FROM tickets WHERE reference_no = ?;