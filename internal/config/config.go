// Package config = configuration loader (C# IOptions + appsettings.json).
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adityax28/fastport/internal/models"
)

// Load reads config.json, then allows env overrides (Render-friendly).
// Env keys mirror ASP.NET style: Portfolio__Email → PORTFOLIO_EMAIL
func Load(path string) (models.PortfolioConfig, error) {
	var cfg models.PortfolioConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	if v := os.Getenv("PORTFOLIO_EMAIL"); v != "" {
		cfg.Email = v
	}
	if v := os.Getenv("PORTFOLIO_GITHUB"); v != "" {
		cfg.GitHub = v
	}
	if v := os.Getenv("PORTFOLIO_LINKEDIN"); v != "" {
		cfg.LinkedIn = v
	}

	return cfg, nil
}

// Environment returns Development / Production (C# IHostEnvironment.EnvironmentName).
func Environment() string {
	if v := os.Getenv("APP_ENV"); v != "" {
		return v
	}
	if v := os.Getenv("GIN_MODE"); v == "release" {
		return "Production"
	}
	return "Development"
}

// Port returns listen port — Render injects PORT; local default 8080.
func Port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}
