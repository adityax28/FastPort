# FastPort — Aditya portfolio (Go + Gin)

Backend portfolio site powered by **Go** and the **Gin** web framework.

## Project layout

| Layer | Folder | Responsibility |
|-------|--------|----------------|
| Entry point | `cmd/server/main.go` | Boot server, wire dependencies, register routes |
| Handlers | `internal/handlers` | HTTP controllers — parse requests, return JSON |
| Services | `internal/service` | Business logic (profile, status, contact validation) |
| Repository | `internal/repository` | Data access (in-memory contact store) |
| Models | `internal/models` | DTOs and JSON shapes |
| Config | `internal/config` + `config.json` | App settings + env overrides |
| Frontend | `web/` | Static HTML, CSS, JS |

## Endpoints

- `GET /api/profile`
- `GET /api/status`
- `POST /api/contact`

## Run locally

```bash
go run ./cmd/server
```

Open http://localhost:8080

## Configure

Edit `config.json`, or set `PORTFOLIO_EMAIL`, `PORTFOLIO_GITHUB`, `PORTFOLIO_LINKEDIN`.

## Deploy (Render)

1. New **Web Service** → connect your repo
2. Language: **Docker**
3. **Root Directory**: `Aditya.Portfolio` (if repo root is the parent folder)
4. **Dockerfile Path**: `./Dockerfile`
5. Deploy — the app listens on `$PORT` (Render injects this automatically)

Optional env vars: `PORTFOLIO_EMAIL`, `PORTFOLIO_GITHUB`, `PORTFOLIO_LINKEDIN`
