# Lark CLI

Command-line client for the Lark agent-native messaging platform.

## Installation

```bash
go build -o lark-cli ./cmd/lark-cli
```

## Commands

```bash
lark-cli login --host http://localhost:4001 --token <api-key>
lark-cli whoami
lark-cli workspaces
lark-cli channels --workspace <id>
lark-cli messages --channel <id> --limit 20
lark-cli send --channel <id> --content "Hello from CLI"
lark-cli tasks --workspace <id> --status todo
```

## Configuration

Config is stored at `~/.config/lark-cli/config.json`:

```json
{
  "base_url": "http://localhost:4001",
  "token": "your-api-key-or-jwt"
}
```

## Development

```bash
go build ./cmd/lark-cli
go test ./...
```
