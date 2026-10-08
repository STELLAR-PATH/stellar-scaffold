<div align="center">

# `stellar-scaffold`

**Standardized CLI Workspace Generator for Soroban Smart Contracts**

[![Stellar Ecosystem](https://img.shields.io/badge/Stellar-Soroban-7B3FE4?style=for-the-badge&logo=stellar)](https://stellar.org)
[![Go 1.22+](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![Drips Stellar Wave](https://img.shields.io/badge/Drips-Stellar%20Wave%20Participant-00D395?style=for-the-badge)](https://drips.network)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

</div>

## 📖 Overview

`stellar-scaffold` is a fast, deterministic scaffolding engine built in Go. It standardizes Soroban contract layouts, sets up unit test suites, and provisions ecosystem configuration boilerplate so developers can start writing business logic immediately.

## ✨ Key Features

- **Standardized Directory Trees**: Clean, maintainable contract folder structures aligned with Stellar best practices.
- **Dependency Pinning**: Automatically locks Soroban SDK versions for reproducible builds.
- **Automated Test Harness**: Generates boilerplate for Rust-based contract integration tests.
- **CI/CD Ready**: Injects GitHub Actions workflows (including `stellarpath-action`) by default.

## 🚀 Installation

Ensure you have Go 1.22+ installed.

```bash
# Clone the repository
git clone https://github.com/STELLAR-PATH/stellar-scaffold.git
cd stellar-scaffold

# Build the CLI
go build -o stellar-scaffold main.go

# (Optional) Move to your PATH
mv stellar-scaffold ~/.local/bin/
```

## 🛠️ Usage

Generate a new, secure Soroban workspace in seconds:

```bash
stellar-scaffold init my-soroban-project
cd my-soroban-project
```

## 🤝 Contributing & Reviewers

**For Contributors:**
- Project templates are stored in the `templates/` directory as Go `text/template` files.
- To add a new template variable, update the `ScaffoldContext` struct in `generator.go`.
- Run `go test ./...` before submitting a PR.

**For Reviewers:**
- Pay attention to the default versions of the Soroban SDK injected into the templates. Ensure they match the current stable Stellar network releases.

---
<div align="center">
  <sub>Part of the <a href="https://github.com/STELLAR-PATH">STELLAR-PATH</a> Toolchain. Built for the Soroban ecosystem.</sub>
</div>
