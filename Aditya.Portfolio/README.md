# Aditya Portfolio

Dynamic portfolio site for **Aditya** — Backend Developer (fintech / .NET).

## Stack

- **ASP.NET Core 8** Minimal APIs (no Razor Pages)
- Plain HTML / CSS / JS in `wwwroot`
- Live endpoints: `/api/profile`, `/api/status`, `/api/contact`

## Run locally

```bash
cd Aditya.Portfolio
dotnet run
```

Open http://localhost:5080

## Deploy on Render

1. New **Web Service** → connect `adityax28/FastPort`
2. Language: **Docker**
3. Branch: **main**
4. **Root Directory**: leave empty
5. **Dockerfile Path**: `./Dockerfile` (or leave default)
6. Plan: **Free**
7. Deploy

Optional env vars: `Portfolio__Email`, `Portfolio__GitHub`, `Portfolio__LinkedIn`

## Configure

Edit `appsettings.json` → `Portfolio` for name, email, GitHub, LinkedIn.
