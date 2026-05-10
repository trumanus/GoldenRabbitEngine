# Volume 2 concept → Go package mapping

The *Golden Rabbit* books are written in Italian prose with mathematical notation. This table links **Volume 2 themes** to **packages** in this repo. It is a **guide for readers and contributors**, not a legal mapping.

| Vol. 2 chapter (theme) | Concept (short) | Package / entrypoints |
|------------------------|-----------------|------------------------|
| 2 | Raw input, encoding, typed event, salience gate, ⊥ | `pkg/event` |
| 3 | Budget **K**, costs κ, feasibility Σκ⪯K | `pkg/budget` |
| 4 | Project tension **E**, PID command **u**, density **d** | `pkg/homeostasis` |
| 5 / 14 | Slow modulators **λ** | `pkg/modulator` |
| 7 | Registry 𝒞_Reg, event schema, ΠΘ, hash | `pkg/registry` |
| 8 | Layers ℓ, short buffer **M_B**, retrieval score | `pkg/memory` (`layers`, `buffer`, `multistore`, `ranking`) |
| 9–10 | Tensions **τ**, conflict resolution | `pkg/values`, `pkg/conflict` |
| 11 | Physical features **z_phys** | `pkg/embodiment` |
| 12 | Partial order, merge | `pkg/distributed` |
| 13 | Tenant isolation guard | `pkg/security` |
| 14 | Regression **η** | `pkg/regression` |
| 14 | Observable identity **Φ** (surrogate) | `pkg/identity` |
| 14 | ACT under cost | `pkg/act` |
| 15 | Outbox pattern | `pkg/outbox` |
| 15 | Forensic append-only chain | `pkg/forensic` |
| 15 | Maturity score **M** | `pkg/maturity` |
| 15 | Orchestrated tick | `pkg/kernel` |
| 2 / 15 | Promotion 𝒥_π, **Applied**, idempotency | `pkg/promotion` |
| — | Deterministic replay steps | `pkg/replay` |
| — | Cognitive snapshot | `pkg/state` |

## Symbol hygiene

Volume 2 uses symbols such as Φ, η, κ, τ, λ with precise meanings. Go identifiers differ by necessity. When extending code, add a short comment `// VOL2: symbol → meaning` at non-obvious boundaries, or extend this document.

## Divergence

When the books and code disagree, **prefer an explicit issue** titled “alignment debt” over silent drift.
