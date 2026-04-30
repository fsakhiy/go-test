// internal/tickets/handler.go
package tickets

import (
	"database/sql"
	"errors"
	"gin-test/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	svc Service
}

func NewHandler(svc Service) *TicketHandler {
	return &TicketHandler{svc: svc}
}

func (h *TicketHandler) Test(c *gin.Context) {
	// Pass the standard request context to your service
	msg, err := h.svc.Test(c.Request.Context())
	if err != nil {
		c.Error(err) // Send to the global ErrorHandler
		c.Abort()    // Stop execution
		return
	}

	// Notice how response.OK doesn't need to be returned anymore
	response.OK(c, 200, "Test successful", msg)
}

func (h *TicketHandler) NewTicket(c *gin.Context) {
	var ticket Ticket

	// ShouldBindJSON automatically maps the JSON and runs validation tags
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	newTicket, err := h.svc.NewTicket(c.Request.Context(), ticket)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, 201, "Ticket created successfully", newTicket)
}

func (h *TicketHandler) GetAll(c *gin.Context) {
	tickets, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, 200, "Tickets fetched successfully", tickets)
}

func (h *TicketHandler) Update(c *gin.Context) {
	// Gin uses Param() instead of Params()
	refNo := c.Param("refNo")
	if refNo == "" {
		c.Error(errors.New("Reference number is required"))
		c.Abort()
		return
	}

	var ticket UpdateTicketRequest

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	updatedTicket, err := h.svc.Update(c.Request.Context(), refNo, ticket)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, 200, "Ticket updated successfully", updatedTicket)
}

func (h *TicketHandler) Delete(c *gin.Context) {
	refNo := c.Param("refNo")
	if refNo == "" {
		c.Error(errors.New("Reference number is required"))
		c.Abort()
		return
	}

	err := h.svc.Delete(c.Request.Context(), refNo)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Err(c, 404, "Ticket not found", err.Error())
			c.Abort()
			return
		}
		c.Error(err)
		c.Abort()
		return
	}

	response.OK[any](c, 200, "Ticket deleted successfully", nil)
}
