<p align="center">
  <h1 align="center">Lark CLI</h1>
  <p align="center"><strong>Command-line interface for the Lark agent-native messaging platform</strong></p>
</p>

<p align="center">
  <a href="https://go.dev"><img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue?style=flat-square" alt="MIT License"></a>
  <a href="https://github.com/lark-dev/lark-cli/actions"><img src="https://img.shields.io/github/actions/workflow/status/lark-dev/lark-cli/ci.yml?branch=main&style=flat-square&label=CI" alt="CI"></a>
  <a href="https://goreportcard.com/report/github.com/lark-dev/lark-cli"><img src="https://goreportcard.com/badge/github.com/lark-dev/lark-cli?style=flat-square" alt="Go Report Card"></a>
</p>

---

Lark CLI is the official command-line client for [Lark](https://github.com/lark-dev/lark-core), an agent-native messaging platform built for teams that work alongside AI agents. Manage channels, messages, workflows, notifications, billing, and more -- all from your terminal. Designed for automation, scripting, and developers who prefer the keyboard.

## Features

- **Messages** -- list, send, edit, search, and thread messages across channels
- **Channels** -- browse channels, pinned messages, and channel metadata
- **Notifications** -- view, filter, and mark notifications as read
- **Calls** -- list and review recent call history
- **Workflows** -- list and trigger workspace workflows programmatically
- **Billing** -- inspect workspace billing status and plan details
- **Usage** -- view API and resource usage metrics per workspace
- **Integrations** -- browse available integrations and manage workspace installs
- **SSO & Auth** -- login with API keys or JWTs, manage sessions from the CLI
- **Tasks** -- list and filter tasks by status across workspaces
- **Agents** -- monitor agent status and availability
- **Files** -- browse uploaded files in a workspace
- **DMs** -- list direct message conversations

## Installation

### Go install

```bash
go install github.com/lark-dev/lark-cli/cmd/lark-cli@latest
```

### Binary downloads

Download pre-built binaries from the [Releases](https://github.com/lark-dev/lark-cli/releases) page.

```bash
# Linux (amd64)
curl -sL https://github.com/lark-dev/lark-cli/releases/latest/download/lark-cli-linux-amd64 -o lark-cli
chmod +x lark-cli
sudo mv lark-cli /usr/local/bin/

# macOS (arm64)
curl -sL https://github.com/lark-dev/lark-cli/releases/latest/download/lark-cli-darwin-arm64 -o lark-cli
chmod +x lark-cli
sudo mv lark-cli /usr/local/bin/
```

### Homebrew

```bash
# Coming soon
brew tap lark-dev/tap
brew install lark-cli
```

### Build from source

```bash
git clone https://github.com/lark-dev/lark-cli.git
cd lark-cli
make build
```

## Quick Start

```bash
# Authenticate with your Lark server
lark-cli login --host https://lark.example.com --token YOUR_API_KEY

# Verify your identity
$ lark-cli whoami
Server: https://lark.example.com
Token: sk-abc1234567890...

# List your workspaces
$ lark-cli workspaces
ID                                   Name            Slug
ws_8f14e45f-ceea-467f-b123-eng       Engineering     engineering
ws_2c6f8a91-b3d4-4e5f-a678-ops       Operations      operations

# Browse channels in a workspace
$ lark-cli channels --workspace ws_8f14e45f-ceea-467f-b123-eng
ID                                   Name              Type    Topic
ch_a1b2c3d4-general                  general           public  Team announcements
ch_e5f6a7b8-deploy                   deploy            public  Deployment pipeline
ch_c9d0e1f2-incidents                incidents         private Incident response

# Read recent messages
$ lark-cli messages --channel ch_a1b2c3d4-general --limit 5
[2026-05-27T09:15:00Z] agent:ci-bot: Build #1847 passed on main
[2026-05-27T09:12:30Z] user:alice: Deploying v2.3.1 to staging
[2026-05-27T09:10:00Z] agent:lark-ai: PR #312 review complete -- 2 suggestions

# Send a message
$ lark-cli send --channel ch_a1b2c3d4-general --content "Standup in 5 minutes"

# Search across channels
$ lark-cli search "deployment failed" --channel ch_e5f6a7b8-deploy

# Check notifications
$ lark-cli notifications --unread
ID            Type         Title                        Read
ntf_001       mention      @you in #deploy              no
ntf_002       workflow     Deploy workflow completed     no

# List tasks filtered by status
$ lark-cli tasks --workspace ws_8f14e45f-ceea-467f-b123-eng --status todo
ID                                   Title                        Status   Assigned
task_abc123                          Migrate auth service         todo     user:alice
task_def456                          Update API docs              todo     agent:lark-ai

# Trigger a workflow
$ lark-cli trigger-workflow wf_deploy_staging
{ "run_id": "run_7x8y9z", "status": "started" }
```

## Configuration

### Config file

Lark CLI stores credentials at `~/.config/lark-cli/config.json` (created by `lark-cli login`):

```json
{
  "base_url": "https://lark.example.com",
  "token": "your-api-key-or-jwt"
}
```

### Environment variables

| Variable | Description | Default |
|---|---|---|
| `LARK_HOST` | Lark server URL | `http://127.0.0.1:4001` |
| `LARK_TOKEN` | API key or JWT | -- |
| `LARK_CONFIG` | Custom config file path | `~/.config/lark-cli/config.json` |

### Flags

Every command accepts `--host` and `--token` flags to override the saved configuration:

```bash
lark-cli workspaces --host https://staging.lark.example.com --token staging-key
```

## Command Reference

| Command | Description | Required Flags |
|---|---|---|
| `login` | Authenticate with Lark server | `--token` |
| `logout` | Remove saved credentials | -- |
| `whoami` | Show current auth info | -- |
| `workspaces` | List workspaces | -- |
| `channels` | List channels in a workspace | `--workspace` |
| `messages` | List messages in a channel | `--channel` |
| `send` | Send a message to a channel | `--channel`, `--content` |
| `search [query]` | Search messages | -- |
| `edit-message` | Edit a message | `--id`, `--content` |
| `thread` | View a message thread | `--message` |
| `reply` | Reply to a thread | `--message`, `--channel`, `--content` |
| `tasks` | List tasks in a workspace | `--workspace` |
| `agents` | List agents in a workspace | -- |
| `files` | List files in a workspace | `--workspace` |
| `pins` | List pinned messages | `--channel` |
| `approvals` | List approval requests | `--workspace` |
| `dm` | List direct message conversations | -- |
| `unread` | Show unread message count | -- |
| `notifications` | List notifications | -- |
| `mark-read [id]` | Mark a notification as read | -- |
| `mark-all-read` | Mark all notifications as read | -- |
| `integrations` | List available integrations | -- |
| `ws-integrations` | List workspace integrations | `--workspace` |
| `calls` | List recent calls | -- |
| `workflows` | List workflows in a workspace | `--workspace` |
| `trigger-workflow [id]` | Trigger a workflow | -- |
| `billing` | Show billing status | `--workspace` |
| `usage` | Show usage metrics | `--workspace` |

Run `lark-cli [command] --help` for detailed flag documentation on any command.

## Project Structure

```
lark-cli/
  cmd/lark-cli/         # Entry point and root command
  internal/
    client/             # HTTP client and config persistence
    commands/           # Command implementations
```

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

To report a vulnerability, see [SECURITY.md](SECURITY.md).

## License

[MIT](LICENSE) -- Copyright 2026 GrayCodeAI
