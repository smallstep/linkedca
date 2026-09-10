# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

## Overview

`linkedca` is a small Go library (module `github.com/smallstep/linkedca`, package `linkedca`) holding the protocol buffer definitions and generated Go code for Linked CAs: the `Majordomo` gRPC service and the messages that describe a CA's configuration, authority, provisioners, admins, ACME accounts, EAB keys, and certificate issuance policy. It is imported by `step-ca` (`github.com/smallstep/certificates`) and `step` (`github.com/smallstep/cli`). There is no binary to run; `make build` is a no-op.

## Commands

```bash
make test               # unit tests with coverage (gotestsum, via go tool)
make race               # unit tests with the race detector
make lint               # golangci-lint (config fetched from smallstep/workflows) + govulncheck
make fmt                # goimports -local github.com/smallstep/linkedca
make generate           # regenerate all *.pb.go from spec/ (see below)
make bootstrap          # go install protoc-gen-go and protoc-gen-go-grpc
make all                # lint + generate + test
```

```bash
go test -run TestAdminFromContext ./...     # run a single test
go build ./...                              # compile check
```

`make test` and `go build ./...` need only a Go toolchain; they complete in a few seconds. `make lint` needs network access (it curls `.golangci.yml` from the `smallstep/workflows` repo). All Go tools (`gotestsum`, `goimports`, `golangci-lint`, `govulncheck`) are `tool` directives in `go.mod`, so nothing needs to be on `$PATH` except `go`, `protoc`, and the two protoc plugins.

CI (`.github/workflows/ci.yml`) calls the shared `smallstep/workflows` `goCI.yml` on every pull request: build, test, lint, govulncheck, and CodeQL across the Go versions in the matrix (`only-latest-golang: false`).

## Generated Code — Do Not Edit

| Pattern | Generator | Source |
|---------|-----------|--------|
| `*.pb.go` (`acme`, `admin`, `config`, `eab`, `majordomo`, `policy`, `provisioners`) | `protoc-gen-go` | `spec/linkedca/<name>.proto` |
| `majordomo_grpc.pb.go` | `protoc-gen-go-grpc` | `spec/linkedca/majordomo.proto` |

Everything else at the repo root is hand-written: `context.go`, `policy.go`, their `_test.go` files, and `tools.go`.

### Regenerating

The `.proto` files under `spec/linkedca/` are the source of truth. To change the API:

1. Edit the relevant `spec/linkedca/*.proto` file.
2. `make bootstrap` once to install `protoc-gen-go` and `protoc-gen-go-grpc` at the versions pinned in `go.mod` (`tools.go` keeps them in the module graph). `protoc` itself must already be installed; it is not managed by this repo.
3. `make generate`. It first checks the plugin versions against the minimums in the Makefile (`GEN_GO_MIN_VERSION`, `GEN_GRPC_MIN_VERSION`), deletes every `*.pb.go`, then runs `protoc --proto_path=spec --go_out=. --go-grpc_out=.` with `--go_opt=module=github.com/smallstep/linkedca` so output lands at the repo root rather than in a nested directory.
4. Commit the regenerated `*.pb.go` alongside the `.proto` change. The header of each generated file records the `protoc` and plugin versions used, so expect a header diff if your local versions differ from the last generation.

There is no `buf` configuration and no `go generate` directive; `make generate` is the only regeneration path. Every proto file uses `package linkedca;` and `option go_package = "github.com/smallstep/linkedca";`, and intra-repo imports are written relative to `spec/` (for example `import "linkedca/policy.proto";`).

## Architecture

```
linkedca/
├── spec/linkedca/
│   ├── majordomo.proto     # Majordomo service: Login, GetConfiguration, provisioner/admin CRUD,
│   │                       #   Post*/Revoke*/Get* certificate RPCs, ACME account CRUD
│   ├── config.proto        # Configuration, Authority, KMS, TLS, SSH, templates (imports admin/provisioners/policy)
│   ├── provisioners.proto  # Provisioner + per-type details (JWK, OIDC, GCP, AWS, Azure, ACME, X5C, K8sSA, SCEP, ...), Claims, Webhook
│   ├── policy.proto        # Policy: X509 / SSH host / SSH user allow-deny name lists
│   ├── admin.proto         # Admin, AdminList
│   ├── acme.proto          # ACMEAccount
│   └── eab.proto           # EABKey (ACME external account binding)
├── *.pb.go                 # generated messages (one per proto), DO NOT EDIT
├── majordomo_grpc.pb.go    # generated MajordomoClient / MajordomoServer, DO NOT EDIT
├── context.go              # hand-written: context helpers for Admin, Provisioner, EABKey
├── policy.go               # hand-written: Policy.Deduplicate()
├── tools.go                # build-tagged imports that pin the protoc plugins in go.mod
└── Makefile
```

The hand-written Go is deliberately thin:

- `context.go` provides `NewContextWith<T>` / `<T>FromContext` / `Must<T>FromContext` triples for `*Admin`, `*Provisioner`, and `*EABKey`, keyed on a private `contextKeyType`. The `FromContext` variants return `ok == false` for a stored nil pointer.
- `policy.go` adds `(*Policy).Deduplicate()`, which removes duplicate entries from every allow/deny name list while preserving order and is nil-safe.

Adding behaviour to a generated type follows the same pattern: put methods in a new hand-written file at the root rather than touching `*.pb.go`.

## Conventions

- Plain Go module, no CLI, no config, no logging, no environment variables.
- Tests use `testify/assert`; `policy_test.go` is table-driven and `context_test.go` uses `t.Parallel()`.
- The repo is public and Apache-2.0 licensed; keep issues, PRs, and commit messages free of anything not appropriate for a public audience.
- Releases are semver tags (`vX.Y.Z`). Downstream consumers pin a tagged version, so a proto change that adds fields is a minor bump; renaming or removing fields or RPCs is a breaking change for every consumer of the generated types.
- Dependabot (`.github/dependabot.yml`) bumps Go modules weekly and `dependabot-auto-merge.yml` auto-merges those PRs; most of the recent history is such bumps.

## Related repos

- `github.com/smallstep/certificates` (`step-ca`) and `github.com/smallstep/cli` (`step`) import this module. To test a proto change against one of them before tagging, add `replace github.com/smallstep/linkedca => ../linkedca` to that repo's `go.mod` locally (do not commit the replace).
- `github.com/smallstep/workflows` supplies the reusable CI workflows and the `.golangci.yml` that `make lint` downloads.
