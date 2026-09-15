// Package models = DTOs / domain shapes (same role as C# records in Models/).
// These structs are serialized to JSON for the frontend — keep camelCase tags stable.
package models

import "time"

// PortfolioConfig maps to appsettings Portfolio section (C# PortfolioOptions).
type PortfolioConfig struct {
	Name            string `json:"name"`
	Title           string `json:"title"`
	Tagline         string `json:"tagline"`
	YearsExperience int    `json:"yearsExperience"`
	Email           string `json:"email"`
	GitHub          string `json:"gitHub"`
	LinkedIn        string `json:"linkedIn"`
	Domain          string `json:"domain"`
}

// ContactRequest = body for POST /api/contact (C# ContactRequest).
// binding tags = Gin validation (similar to DataAnnotations).
type ContactRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required,min=10"`
}

// ContactResponse = API response for contact form.
type ContactResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// StoredContact = what the repository persists in memory.
type StoredContact struct {
	At      time.Time      `json:"at"`
	Request ContactRequest `json:"request"`
}

// ProjectItem = one work item on the portfolio.
type ProjectItem struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
	Impact  string   `json:"impact"`
}

// ProfilePayload = GET /api/profile response (must match site.js field names).
type ProfilePayload struct {
	Name            string        `json:"name"`
	Title           string        `json:"title"`
	Tagline         string        `json:"tagline"`
	YearsExperience int           `json:"yearsExperience"`
	Domain          string        `json:"domain"`
	Email           string        `json:"email"`
	GitHub          string        `json:"gitHub"`
	LinkedIn        string        `json:"linkedIn"`
	Focus           []string      `json:"focus"`
	Projects        []ProjectItem `json:"projects"`
	Stack           []string      `json:"stack"`
}

// SystemStatus = GET /api/status response (used by UI + UptimeRobot keep-alive).
type SystemStatus struct {
	Status        string    `json:"status"`
	Environment   string    `json:"environment"`
	StartedAt     time.Time `json:"startedAt"`
	UptimeSeconds float64   `json:"uptimeSeconds"`
	Runtime       string    `json:"runtime"`
	Pipelines     []string  `json:"pipelines"`
}
