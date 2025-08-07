# Base stage for common setup
FROM golang:1.21.6 AS base
ENV PORT=8080
COPY . /src/

# Development build stage
FROM base AS development
# Create binary with no dependencies on C libraries
ENV CGO_ENABLED=0
WORKDIR /src
RUN go build -o /app/esync ./cmd/esync/main.go
COPY web/static/index.html /app/web/static/index.html
EXPOSE 8080
CMD ["/app/esync"]

# Debug build stage
FROM base AS debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /src
RUN go build -gcflags="all=-N -l" -o /app/esync ./cmd/esync/main.go
RUN cp $GOPATH/bin/dlv /app/

EXPOSE 8080 2345
CMD ["/app/dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/esync"]

# Production build stage
FROM debian:stable-slim AS production
COPY --from=development /app/esync /app/esync
EXPOSE 8080
CMD ["/app/esync"]