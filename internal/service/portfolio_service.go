// Package service = business logic layer (C# PortfolioData + contact rules).
// Handlers call services; services call repositories. No Gin types here.
package service

import (
	"runtime"
	"strings"
	"time"

	"github.com/adityax28/fastport/internal/models"
	"github.com/adityax28/fastport/internal/repository"
)

// PortfolioService builds profile/status payloads (C# PortfolioData).
type PortfolioService struct {
	cfg       models.PortfolioConfig
	startedAt time.Time
}

// NewPortfolioService wires config (like injecting IOptions<PortfolioOptions>).
func NewPortfolioService(cfg models.PortfolioConfig) *PortfolioService {
	return &PortfolioService{
		cfg:       cfg,
		startedAt: time.Now().UTC(),
	}
}

// GetProfile = domain logic for GET /api/profile.
func (s *PortfolioService) GetProfile() models.ProfilePayload {
	return models.ProfilePayload{
		Name:            s.cfg.Name,
		Title:           s.cfg.Title,
		Tagline:         s.cfg.Tagline,
		YearsExperience: s.cfg.YearsExperience,
		Domain:          s.cfg.Domain,
		Email:           s.cfg.Email,
		GitHub:          s.cfg.GitHub,
		LinkedIn:        s.cfg.LinkedIn,
		Focus: []string{
			"Transaction & payment systems",
			"API integrations at production scale",
			"Database & performance optimization",
			"Owning bottlenecks end-to-end",
		},
		Projects: []models.ProjectItem{
			{
				ID:      "pay-core",
				Title:   "Payment transaction pipeline",
				Summary: "Backend flows for payment initiation, status tracking, and reconciliation-oriented processing in a fintech environment.",
				Tags:    []string{"C#", ".NET", "SQL", "APIs"},
				Impact:  "Reliable handling of high-value money movement paths.",
			},
			{
				ID:      "api-mesh",
				Title:   "Upstream / downstream integrations",
				Summary: "Hardened API integrations with external and internal services—timeouts, retries, and clear failure surfaces for operators.",
				Tags:    []string{".NET", "HTTP", "CI/CD"},
				Impact:  "Fewer silent failures; faster incident diagnosis.",
			},
			{
				ID:      "db-perf",
				Title:   "Query & throughput optimization",
				Summary: "Identified hot paths, tightened queries and indexes, and reduced latency under production load.",
				Tags:    []string{"SQL Server", "Profiling", "Caching"},
				Impact:  "Measurable wins on critical transaction endpoints.",
			},
			{
				ID:      "ops-ownership",
				Title:   "Production problem solving",
				Summary: "Took ownership beyond tickets: mapped business impact, isolated bottlenecks, and shipped durable fixes.",
				Tags:    []string{"Observability", "Incident response"},
				Impact:  "From firefighting to scalable systems thinking.",
			},
		},
		Stack: []string{
			"C#", ".NET", "ASP.NET Core", "SQL Server", "REST APIs",
			"CI/CD", "Git", "Caching", "Logging", "Fintech domain", "Go", "Gin",
		},
	}
}

// GetStatus = domain logic for GET /api/status.
func (s *PortfolioService) GetStatus(environment string) models.SystemStatus {
	return models.SystemStatus{
		Status:        "operational",
		Environment:   environment,
		StartedAt:     s.startedAt,
		UptimeSeconds: time.Since(s.startedAt).Seconds(),
		Runtime:       runtime.Version(),
		Pipelines:     []string{"build", "test", "deploy"},
	}
}

// ContactService = use-case for submitting contact messages (validation + repo).
type ContactService struct {
	repo *repository.ContactRepository
}

// NewContactService injects the repository (constructor DI).
func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

// Submit validates/normalizes and stores the message.
// Returns (response, ok). ok=false means client error.
func (s *ContactService) Submit(req models.ContactRequest) (models.ContactResponse, bool) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	// Extra guard beyond Gin binding (mirrors Program.cs checks).
	if len(req.Name) < 2 || req.Email == "" || len(req.Message) < 10 {
		return models.ContactResponse{
			Ok:      false,
			Message: "Name, valid email, and message (10+ chars) are required.",
		}, false
	}

	s.repo.Add(req)

	// Production: send email / enqueue here.
	return models.ContactResponse{
		Ok:      true,
		Message: "Message received. I'll get back to you soon.",
	}, true
}
