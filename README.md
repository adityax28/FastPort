# FastPort — Aditya portfolio (Go + Gin)

Learning-oriented layout (ASP.NET → Gin mapping):

| Layer | Folder | ASP.NET equivalent |
|-------|--------|--------------------|
| Entry / DI | `cmd/server/main.go` | `Program.cs` |
| Controllers | `internal/handlers` | Controllers / Minimal API handlers |
| Services | `internal/service` | `Services/` business logic |
| Repository | `internal/repository` | data access (`ContactStore`) |
| Models | `internal/models` | DTOs / records |
| Config | `internal/config` + `config.json` | `appsettings.json` |
| Frontend | `web/` | `wwwroot/` |

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

Docker runtime, Dockerfile at repo root. Listens on `$PORT`.
