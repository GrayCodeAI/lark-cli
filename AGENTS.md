# CLAUDE.md — lark-cli

## Build & Test

```bash
go build ./cmd/lark-cli
go test ./... -race
```

## Structure

- `cmd/lark-cli/main.go` — Entry point, subcommand dispatch
- `internal/client/client.go` — HTTP client, config persistence
- `internal/commands/commands.go` — Command implementations

## Conventions

- Config stored at `~/.config/lark-cli/config.json`
- API token passed via `--token` flag or persisted config
- All output via `text/tabwriter` for tabular formatting
- Uses `cobra` for CLI framework
