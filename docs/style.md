# Code Style & Conventions

This document summarizes coding style, conventions and recommendations observed in this project and provides actionable suggestions to keep the codebase consistent and maintainable.

Scope
- Files considered: `cmd/api/main.go`, `internal/domain/*`, `internal/infra/*`, `internal/handlers/*`, `internal/middleware/*`, `internal/service/*`, `internal/ports/*`, `docs/*`.

1. Project layout
- Current: follows a Clean/Ports-and-Adapters hybrid. Top-level `cmd/`, `internal/` and `docs/` folders.
- Recommendation: keep this layout; continue to place application entrypoints in `cmd/`, business logic in `internal/service`, domain entities in `internal/domain`, adapters in `internal/infra` and HTTP handlers in `internal/handlers`.

2. Packages & naming
- Use short, lowercase package names (e.g. `infra`, `service`, `handlers`, `domain`).
- Export only what is required. Keep fields unexported by default.
- Recommendation: prefer `repository` or `repo` only if clearer than `infra` for DB-specific code.

3. Dependency injection
- Current: services are constructed in `cmd/api/main.go` and injected into handlers (good).
- Recommendation: keep wiring in `cmd/` only. Do not call infra or DB directly from `service` or `handlers`—use interfaces (ports).

4. Interfaces (ports)
- Current: `internal/ports/user_repo.go` defines a `UserRepository` interface used by `service`.
- Recommendation: keep interfaces small and purpose-specific. Document behaviour (e.g. return `ErrNotFound` vs `gorm.ErrRecordNotFound`).

5. Domain models vs persistence models
- Current: `internal/domain.User` contains `gorm` and `json` tags (mixed concerns).
- Recommendation: if you expect growth, consider splitting domain entity (pure) from persistence model + mapper. For small projects, current approach is acceptable.

6. Error handling
- Use sentinel errors for expected error cases (e.g. `ErrUserExists`) and wrap with context when necessary.
- Avoid exposing internal DB errors directly to HTTP responses; map them to user-friendly messages.

7. Logging
- Use `log` for simple cases. For production, prefer structured logger (e.g. `zap` or `logrus`).
- Include request correlation id in logs for important flows (transfer/purchase later).

8. Security & secrets
- Current: JWT secret is hard-coded (`secret_key_change_me`).
- Must: move secrets and config (DB path, JWT secret, port) to environment variables or a config provider.

9. HTTP handlers
- Keep handlers thin: validate input, call service, map response/HTTP codes.
- Current handlers mostly follow this. Continue using Gin binding for validation.

10. JWT middleware
- Current: middleware parses token and sets `user_id` in context. Good pattern.
- Recommendation: return typed context values (helper to extract `UserID(c)`), and validate token error types to provide clearer responses.

11. Tests
- Current: unit tests use `testify/mock` for service layer. Good separation.
- Recommendation:
  - Add handler tests (httptest + mocked service).
  - Add integration tests using SQLite in-memory `file::memory:?cache=shared` for repository tests.
  - Tag integration tests (build tag `integration`) so CI can run unit tests fast.

12. Formatting & linting
- Use `gofmt`/`gofumpt` for formatting.
- Add `golangci-lint` to CI with rules enabled (errcheck, govet, staticcheck, gosimple, gofmt).

13. Modules & dependencies
- Keep `go.mod` tidy. Prefer running `go mod tidy` after dependency changes.
- Vendor only if required by your deployment strategy.

14. Swagger / API docs
- API docs live in `docs/` (OpenAPI + generated files). Keep docs updated when handlers change.
- Recommendation: add `swag` comment annotations on handlers for automated generation and include `swag init` in developer setup instructions.

15. CI / Developer commands (recommended)
- Add `Makefile` with targets:
  - `make build` — build binary
  - `make run` — run app locally
  - `make test` — run unit tests
  - `make itest` — run integration tests
  - `make lint` — run linters
  - `make tidy` — run `go mod tidy`

16. Misc recommendations
- Use context (`context.Context`) in service methods to support cancellation and tracing.
- Add API versioning prefix (e.g. `/v1`) in routes in the future.
- Add OpenAPI examples for error responses and auth header in docs.

17. Quick checklist for PRs
- gofmt / go vet passed
- tests added (unit tests for logic, integration for DB-critical flows)
- swagger updated if API changes
- secrets not committed

---

If you want, I can:
- Create a `Makefile` and GitHub Actions workflow for lint/test.
- Add a small helper `config` package to read env vars and replace hard-coded values.
- Split `domain` and persistence models and add mappers.

Which follow-up would you like? (Makefile / config package / split models / add linter CI)
