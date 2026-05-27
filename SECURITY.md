# Security Policy

## Supported Versions

| Version | Supported |
|---|---|
| Latest release | Yes |
| Older releases | No |

## Reporting a Vulnerability

If you discover a security vulnerability in Lark CLI, please report it responsibly.

**Do not open a public GitHub issue for security vulnerabilities.**

Instead, email: **security@graycodeai.com**

Include the following in your report:

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

## What to Expect

- **Acknowledgment** within 48 hours of your report
- **Status update** within 7 days with an initial assessment
- **Resolution target** within 30 days for confirmed vulnerabilities
- Credit in the release notes (unless you prefer to remain anonymous)

## Scope

The following are in scope:

- Authentication and token handling in `lark-cli`
- Credential storage (`~/.config/lark-cli/config.json`)
- HTTP client behavior (TLS, header injection, SSRF)
- Dependency vulnerabilities in Go modules

The following are out of scope:

- The Lark server itself (report separately to [lark-core](https://github.com/lark-dev/lark-core))
- Social engineering attacks
- Denial of service against the CLI itself

## Best Practices for Users

- Never commit your `config.json` or API tokens to version control
- Use short-lived tokens when possible
- Run `lark-cli logout` when finished on shared machines
- Keep `lark-cli` updated to the latest release
