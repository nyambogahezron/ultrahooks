# Security Policy

## Supported Versions

Currently, only the latest major release of UltraHooks receives continuous automated security updates and patches. 

| Version | Supported          |
| ------- | ------------------ |
| v1.0.x  | :white_check_mark: |
| < v1.0  | :x:                |

## Reporting a Vulnerability

We take the security of UltraHooks extremely seriously. UltraHooks runs directly inside a user's development machine during Git hooks, giving it access to execute system shell commands. It is imperative that vulnerabilities allowing command injection or unauthorized execution bypass are handled immediately.

If you discover a security vulnerability, please **DO NOT** open a public GitHub issue. 

Instead, please email the maintainers directly at [YOUR_EMAIL@DOMAIN.com] with the subject `[SECURITY] UltraHooks Vulnerability`.

Please include the following in your report:
- Type of issue (e.g. Command Injection, Buffer Overflow, File System Escape, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- A proof-of-concept repository or exact steps to reproduce the issue reliably
- The impact of the issue on the end-user

We will triage the vulnerability within 48 hours and work with you to release a patched version safely.
