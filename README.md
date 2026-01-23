# mpp-project

## Burden of proof 5.1 (Bewijslast)

### Level 1 — Unit tests locally
Run unit tests from your IDE or terminal:
```bash
make test
# or
go test ./...
```

### Level 2 — Unit tests + linter in the development pipeline
A GitHub Actions pipeline is included at `.github/workflows/ci.yml` that runs on every push/PR:
- `go test ./... -race`
- `golangci-lint run` (configured via `.golangci.yml`)

Run locally:
```bash
make ci
```

### Level 2+ — Advanced static analysis + metrics
`golangci-lint` runs multiple analyzers that report code smells and complexity (e.g. `staticcheck`, `gocritic`, `revive`, `gocyclo`).
Coverage metrics are produced with:
```bash
make coverage
# optional HTML report:
make coverage-html
```

### Level 3 — Coverage increases every sprint (proven by metric output)
A coverage gate script compares the current total coverage against:
- a minimum (`COVERAGE_MIN`, default 0), and
- a baseline stored in `.coverage_baseline`.

In each sprint you can update `.coverage_baseline` to the new achieved total coverage, so the next sprint must exceed it.

```bash
# prints the total coverage line (proof)
make coverage

# enforce minimum + 'must be higher than baseline'
COVERAGE_MIN=0 make coverage-gate
```
