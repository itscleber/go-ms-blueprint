# Go Microservice Blueprint

[![CI](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci-pr.yaml/badge.svg)](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci-pr.yaml/badge.svg)

A lightweight and opinionated Go microservice blueprint with structured folders, built-in observability, OpenTelemetry integration, and Docker support.

---

## 🗂️ Project Structure

```
.
├── Dockerfile
├── Makefile
├── cmd/
│   └── main.go
├── infra/
│   └── otel/
│       └── otel-collector-config.yaml
├── internal/
│   ├── api/
│   │   ├── health_routes.go
│   │   ├── ops_routes.go
│   │   ├── router.go
│   │   └── routes.go
│   ├── config/
│   │   ├── bootstrap.go
│   │   └── logger.go
│   ├── handlers/
│   │   ├── health.go
│   │   ├── liveness.go
│   │   └── readiness.go
│   ├── repositories/
│   │   ├── health.go
│   │   ├── liveness.go
│   │   └── readiness.go
│   ├── services/
│   │   ├── health.go
│   │   ├── liveness.go
│   │   └── readiness.go
│   └── telemetry/
│       └── tracer.go
├── tests/
│   └── health_test.go
├── docker-compose.yaml
├── docker-compose.override.yaml
├── go.mod
├── go.sum
└── .env.example
```

---

## 🧪 Makefile Targets

- `make up` – Start services with Docker Compose (local mode)
- `make run` – Run the service via `go run`
- `make test` – Run unit tests inside the container
- `make lint` – Run `golangci-lint` over the codebase
- `make cover` – Generate code coverage report
- `make down` – Tear down Docker Compose environment
- `make pre-commit-install` – Install Git pre-commit hooks
- `make pre-commit` – Run all pre-commit hooks manually

---

## ⚙️ Local Development

Start the service locally:

```sh
make up
```

The service will be available at: `http://localhost:8080`
Port can be customized via the `.env` file.

Health and ops endpoints:

- `GET /v1/health`
- `GET /ops/ready`
- `GET /ops/live`

---

## 🧵 Observability (OpenTelemetry)

Traces are sent from the application to the local OpenTelemetry Collector, which then exports to Jaeger.

- OTEL Collector config: `infra/otel/otel-collector-config.yaml`
- Jaeger UI: [http://localhost:16686](http://localhost:16686)

---

## ✅ GitHub Actions CI

CI pipeline runs on every PR and push to `main`, performing:

[CI Workflow Status](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci.yaml)

---

## 🚀 Getting Started with the Boilerplate

To use this blueprint as a starting point for your own microservice:

```bash
# Clone without Git history
git clone --depth=1 https://github.com/itscleber/go-ms-blueprint my-new-service
cd my-new-service
rm -rf .git

# Initialize your own Git repository
git init
git remote add origin <your-repo-url>

# Update the Go module name
go mod edit -module github.com/your-org/your-service-name
go mod tidy
```

You can now customize the service name, API routes, handlers, and telemetry setup according to your needs.

### 🐳 Service-Specific Docker Configuration

Use `docker-compose.override.yaml` to add service-specific configurations such as:
- Additional environment variables
- Volume mounts for local development
- Service dependencies (databases, caches, etc.)
- Custom ports or network configurations

This file is automatically loaded by Docker Compose and allows you to extend the base configuration without modifying the main `docker-compose.yaml` file.
