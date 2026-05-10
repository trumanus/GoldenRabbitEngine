# golden-rabbit-engine

**Go reference library** inspired by the *Golden Rabbit* Volume 2 narrative: typed events, constitutional **Registry**, cognitive budget **K**, homeostasis, promotion with idempotency, multi-layer memory, and related building blocks.

This repository is **indicative reference code** for engineers and auditors: it encodes concepts as testable APIs. It is **not** a turnkey production stack (no LLM runtime, message broker, encrypted signing, or SQL persistence in-tree).

## Maintainer

**Andrea Lagomarsini** (@trumanus) — [trumanus@gmail.com](mailto:trumanus@gmail.com). See `AUTHORS`.

Italian-language overview: [`README.it.md`](README.it.md).

## Requirements

- Go 1.22 or newer

## Quick start

```bash
git clone https://github.com/trumanus/GoldenRabbitEngine.git
cd GoldenRabbitEngine
go test ./...
go run ./cmd/gr-demo
go run ./cmd/gr-full
```

### Use as a module

```go
import "github.com/trumanus/GoldenRabbitEngine/pkg/registry"
```

(Published tags follow SemVer, e.g. `v0.1.0`.)

## Documentation

| Document | Purpose |
|----------|---------|
| [docs/DISCLAIMER.md](docs/DISCLAIMER.md) | Scope: reference vs production |
| [docs/VOLUME2_MAPPING.md](docs/VOLUME2_MAPPING.md) | Book concepts ↔ packages |
| [CONTRIBUTING.md](CONTRIBUTING.md) | How to contribute |
| [SECURITY.md](SECURITY.md) | Reporting issues |
| [CHANGELOG.md](CHANGELOG.md) | Release notes |
| [RELEASING.md](RELEASING.md) | Tags and releases |

## Package layout (summary)

| Area | Package |
|------|---------|
| Registry & constitution | `pkg/registry` |
| Event chain x→z→e | `pkg/event` |
| Budget **K** | `pkg/budget` |
| Homeostasis | `pkg/homeostasis` |
| Slow modulators **λ** | `pkg/modulator` |
| Memory layers & ranking | `pkg/memory` |
| Value tensions **τ**, conflict | `pkg/values`, `pkg/conflict` |
| ACT / cost | `pkg/act` |
| Embodiment | `pkg/embodiment` |
| Causal ordering | `pkg/distributed` |
| Tenant isolation checks | `pkg/security` |
| Regression **η** | `pkg/regression` |
| Observable identity **Φ** | `pkg/identity` |
| Promotion **π**, idempotency | `pkg/promotion` |
| Replay | `pkg/replay` |
| State / snapshots | `pkg/state` |
| Outbox | `pkg/outbox` |
| Forensic append-only log | `pkg/forensic` |
| Maturity score **M** | `pkg/maturity` |
| Tick orchestration | `pkg/kernel` |

Details and chapter references: **docs/VOLUME2_MAPPING.md**.

## License

MIT — see [LICENSE](LICENSE).

## Relationship to the books

The *Golden Rabbit* books are the conceptual specification; this repo is an **optional implementation sketch**. When they diverge, treat mismatches as **alignment debt** and track them in issues (see Volume 2, implementation chapter).
