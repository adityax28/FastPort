# Build (Go + Gin — replaces .NET SDK image)
FROM golang:1.22-alpine AS build
WORKDIR /src

COPY go.mod ./
COPY go.sum* ./
RUN go mod download

COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o /out/server ./cmd/server

# Runtime
FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache ca-certificates

ENV GIN_MODE=release
ENV APP_ENV=Production

COPY --from=build /out/server /app/server
COPY config.json /app/config.json
COPY web /app/web

EXPOSE 8080

# Render injects PORT; our app reads os.Getenv("PORT")
CMD ["/app/server"]
