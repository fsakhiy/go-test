package tickets

type UpdateTicketRequest struct {
	// Here, user_id and reference_no are omitted because you shouldn't
	// be able to change them during an update!

	// validate:"omitempty" means if they don't send it, that's fine.
	// But if they DO send it, it must match the 'oneof' rule.
	Status string `json:"status" validate:"required"`
}
