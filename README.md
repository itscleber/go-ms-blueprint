# Go Microservice Blueprint

[![CI](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci-pr.yaml/badge.svg)](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci-pr.yaml/badge.svg)

A lightweight and opinionated Go microservice blueprint with structured folders, built-in observability, and Docker support.

---

## 🗂️ Project Structure

```
.
├── Dockerfile
├── Makefile
├── cmd
│   └── main.go
├── internal
│   ├── api
│   │   ├── routes.go
│   │   └── health_routes.go
│   ├── handlers
│   │   └── health.go
│   ├── services
│   │   └── health.go
│   ├── repositories
│   │   └── health.go
│   ├── config
│   │   └── logger.go
│   └── telemetry
│       └── tracer.go
├── docker-compose.dev.yaml
├── go.mod
├── go.sum
└── tests
    └── health_test.go
```

---

## 🛠️ Usage

All commands are available through `make`:

- `make run` — Run the application locally
- `make test` — Run unit tests
- `make pre-commit-install` — Install Git pre-commit hooks
- `make pre-commit-run` — Run all pre-commit hooks manually

The service listens on the port specified by the `PORT` environment variable.
If not set, it defaults to `8080`.

---

## 🔍 Jaeger Tracing

This project includes OpenTelemetry integration with traces sent to Jaeger.
Default endpoint: `http://localhost:16686`

Access the Jaeger UI at `http://localhost:16686`

---

## ✅ GitHub Actions CI

GitHub Actions automatically lints and tests your code on every push and pull request to `main`.

[CI Workflow Status](https://github.com/itscleber/go-ms-blueprint/actions/workflows/ci.yaml)
