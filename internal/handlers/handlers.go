// Package handlers = CONTROLLERS in ASP.NET terms.
// They parse HTTP (Gin context), call services, return JSON.
// No business rules / no DB access here — only request/response wiring.
package handlers

import (
	"net/http"

	"github.com/adityax28/fastport/internal/config"
	"github.com/adityax28/fastport/internal/models"
	"github.com/adityax28/fastport/internal/service"
	"github.com/gin-gonic/gin"
)

// PortfolioHandler groups profile + status endpoints (like a PortfolioController).
type PortfolioHandler struct {
	portfolio *service.PortfolioService
}

// NewPortfolioHandler = controller constructor.
func NewPortfolioHandler(portfolio *service.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{portfolio: portfolio}
}

// GetProfile — ENDPOINT: GET /api/profile
// Maps to: api.MapGet("/profile", ...)
func (h *PortfolioHandler) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, h.portfolio.GetProfile())
}

// GetStatus — ENDPOINT: GET /api/status
// Maps to: api.MapGet("/status", ...)
func (h *PortfolioHandler) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, h.portfolio.GetStatus(config.Environment()))
}

// ContactHandler = ContactController.
type ContactHandler struct {
	contacts *service.ContactService
}

// NewContactHandler = controller constructor.
func NewContactHandler(contacts *service.ContactService) *ContactHandler {
	return &ContactHandler{contacts: contacts}
}

// PostContact — ENDPOINT: POST /api/contact
// Maps to: api.MapPost("/contact", ...)
func (h *ContactHandler) PostContact(c *gin.Context) {
	var req models.ContactRequest

	// ShouldBindJSON = model binding + validation (binding tags on the struct).
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ContactResponse{
			Ok:      false,
			Message: "Name is required. A valid email is required. Message should be at least 10 characters.",
		})
		return
	}

	resp, ok := h.contacts.Submit(req)
	if !ok {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
