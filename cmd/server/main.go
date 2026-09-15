// Package main = application entry + composition root (C# Program.cs).
// Here we: load config → create repo/services/handlers → register ROUTES (endpoints) → serve static UI.
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/adityax28/fastport/internal/config"
	"github.com/adityax28/fastport/internal/handlers"
	"github.com/adityax28/fastport/internal/repository"
	"github.com/adityax28/fastport/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// --- Composition root (manual DI) ---
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// REPOSITORY layer
	contactRepo := repository.NewContactRepository()

	// SERVICE layer
	portfolioSvc := service.NewPortfolioService(cfg)
	contactSvc := service.NewContactService(contactRepo)

	// CONTROLLER / HANDLER layer
	portfolioHandler := handlers.NewPortfolioHandler(portfolioSvc)
	contactHandler := handlers.NewContactHandler(contactSvc)

	if config.Environment() == "Production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// ROUTER (Gin engine) — like WebApplication + MapGroup
	r := gin.Default()

	// --- API ENDPOINTS (controllers wired to routes) ---
	api := r.Group("/api")
	{
		// ENDPOINT GET /api/profile
		api.GET("/profile", portfolioHandler.GetProfile)
		// ENDPOINT GET /api/status
		api.GET("/status", portfolioHandler.GetStatus)
		// ENDPOINT POST /api/contact
		api.POST("/contact", contactHandler.PostContact)
	}

	// --- Static frontend (was wwwroot in ASP.NET) ---
	webRoot := resolveWebRoot()
	r.Static("/css", filepath.Join(webRoot, "css"))
	r.Static("/js", filepath.Join(webRoot, "js"))
	r.StaticFile("/", filepath.Join(webRoot, "index.html"))
	r.StaticFile("/index.html", filepath.Join(webRoot, "index.html"))

	// SPA-style fallback for unknown non-API paths → index.html
	r.NoRoute(func(c *gin.Context) {
		if len(c.Request.URL.Path) >= 4 && c.Request.URL.Path[:4] == "/api" {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			return
		}
		c.File(filepath.Join(webRoot, "index.html"))
	})

	addr := ":" + config.Port()
	log.Printf("FastPort (Gin) listening on %s — web=%s", addr, webRoot)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func resolveWebRoot() string {
	candidates := []string{"web", "./web", "../web"}
	for _, c := range candidates {
		if st, err := os.Stat(filepath.Join(c, "index.html")); err == nil && !st.IsDir() {
			abs, _ := filepath.Abs(c)
			return abs
		}
	}
	return "web"
}
