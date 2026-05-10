# Contributing

Thank you for improving **golden-rabbit-engine**. This project is a **reference library**: clarity, determinism where promised, and honest documentation beat feature volume.

## Principles

1. **Match the book’s intent** — When adding behavior, cite which Volume 2 theme it implements (comment or PR description).
2. **Small PRs** — One conceptual concern per pull request when possible.
3. **Tests** — New logic should include `go test` coverage; simulation/replay tests are especially welcome.
4. **English** — Issue titles, PR descriptions, and user-facing docs in this repo use **English**. Code comments may be Italian only if duplicated with a one-line English summary for maintainers.

## Development

```bash
go test ./...
go vet ./...
```

Optional formatting:

```bash
gofmt -w .
```

## Governance

- **Maintainer:** Andrea Lagomarsini (trumanus) — see `AUTHORS`.
- **License:** MIT — contributions are accepted under the same license unless explicitly stated otherwise.

## Code of conduct

See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
