# Releasing

## Versioning

Use **Semantic Versioning** tags on GitHub:

```bash
git tag -a v0.1.0 -m "v0.1.0 reference implementation"
git push origin v0.1.0
```

## Pre-flight

```bash
go test ./...
go vet ./...
```

## Go module path

The module path is `github.com/trumanus/GoldenRabbitEngine`. If you fork or move the repository, update `go.mod` and all internal imports consistently.

## GitHub Release

Create a **Release** from the tag and attach optional checksums or SBOM if your organisation requires them.
