# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-05-27

### Added

- Initial release of `lark-cli`
- Authentication: `login`, `logout`, `whoami`
- Workspace and channel browsing: `workspaces`, `channels`
- Messaging: `messages`, `send`, `edit-message`, `search`, `thread`, `reply`
- Direct messages: `dm`, `unread`
- Notifications: `notifications`, `mark-read`, `mark-all-read`
- Task management: `tasks` with status filtering
- Agent monitoring: `agents`
- File browsing: `files`
- Pinned messages: `pins`
- Approvals: `approvals`
- Integrations: `integrations`, `ws-integrations`
- Calls: `calls` with limit support
- Workflows: `workflows`, `trigger-workflow`
- Billing and usage: `billing`, `usage`
- Persistent config stored at `~/.config/lark-cli/config.json`
- Tabular output via `text/tabwriter`
- CI pipeline with build, test, and vet
