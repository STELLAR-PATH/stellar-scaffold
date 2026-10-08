<div align="center">

<h1><code>stellar-scaffold</code></h1>
<h3>Standardized CLI Workspace Generator for Soroban Smart Contracts</h3>

[![Stellar Ecosystem](https://img.shields.io/badge/Stellar-Soroban-7B3FE4?style=for-the-badge&logo=stellar)](https://stellar.org)
[![Rust 2021](https://img.shields.io/badge/Rust-2021-DEA584?style=for-the-badge&logo=rust)](https://www.rust-lang.org)
[![Drips Stellar Wave](https://img.shields.io/badge/Drips-Stellar%20Wave%20Participant-00D395?style=for-the-badge)](https://drips.network)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

</div>

---

## 1. Executive Summary

`stellar-scaffold` is a robust, highly-deterministic scaffolding generator written in Go. Its primary directive is to eliminate configuration drift across the Soroban ecosystem by generating mathematically precise, standardized workspace topologies.

Taking cues from ecosystem pioneers like **SoroTrail**, this engine strictly provisions directories, rust toolchains, Makefile automation, and integration test harnesses right out of the box, allowing developers to focus purely on contract business logic.

---

## 2. Core Architecture & Templating

### Go `text/template` Engine
The core generator parses a strictly defined schema from embedded Go templates. The architecture prevents syntax errors in generated Rust code by running a localized `cargo check` validation immediately after generation.


```text
       +-------------------------------------------------------------+
       |                  stellar-scaffold (Go CLI)                  |
       |  * Standardized directory trees                             |
       |  * Soroban SDK dependency pinning                           |
       |  * Automated test harness boilerplate                       |
       |  * Integrated GitHub Actions CI pipelines                   |
       +-------------------------------------------------------------+
```


---

## 3. Installation Specifications

Ensure you have **Go 1.22+** installed on your system.

### Method A: Direct Go Install

- `go⠀install⠀github.com/STELLAR-PATH/stellar-scaffold@latest`


### Method B: Source Build

- `git⠀clone⠀https://github.com/STELLAR-PATH/stellar-scaffold.git`
- `cd⠀stellar-scaffold`
- `go⠀build⠀-o⠀stellar-scaffold⠀main.go`
- `sudo⠀mv⠀stellar-scaffold⠀/usr/local/bin/`


---

## 4. CLI Command Matrix

| Command | Flags / Arguments | Description | Example |
| :--- | :--- | :--- | :--- |
| `init` | `<PROJECT_NAME> [--no-git]` | Generates the complete Soroban workspace directory topology. | `stellar-scaffold init defi-pool` |
| `add` | `contract <NAME>` | Injects a new secondary contract into an existing workspace. | `stellar-scaffold add contract oracle` |
| `upgrade` | `--latest` | Bumps the Soroban SDK dependencies to the latest stable Stellar network release. | `stellar-scaffold upgrade` |
| `ci` | `--provider [github, gitlab]`| Injects CI pipelines (including `stellarpath-action`) into the repository. | `stellar-scaffold ci --provider github` |

---

## 5. Target Template Topology

When `stellar-scaffold init` is invoked, it deterministically constructs the following file system layout:


- `my-soroban-project/`
- `├──⠀.github/`
- `│⠀⠀⠀└──⠀workflows/`
- `│⠀⠀⠀⠀⠀⠀⠀└──⠀stellarpath.yml⠀⠀⠀⠀⠀⠀#⠀Automated⠀AST⠀PR⠀gatekeeper`
- `├──⠀contracts/`
- `│⠀⠀⠀└──⠀main_contract/`
- `│⠀⠀⠀⠀⠀⠀⠀├──⠀src/`
- `│⠀⠀⠀⠀⠀⠀⠀│⠀⠀⠀├──⠀lib.rs⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀#⠀Core⠀contract⠀logic`
- `│⠀⠀⠀⠀⠀⠀⠀│⠀⠀⠀├──⠀storage.rs⠀⠀⠀⠀⠀⠀⠀#⠀Secure⠀DataKey⠀enums`
- `│⠀⠀⠀⠀⠀⠀⠀│⠀⠀⠀└──⠀test.rs⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀#⠀Integration⠀test⠀harness`
- `│⠀⠀⠀⠀⠀⠀⠀└──⠀Cargo.toml⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀#⠀Pinned⠀SDK⠀versions`
- `├──⠀Makefile⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀#⠀Build⠀&⠀Deployment⠀automation`
- `├──⠀.cargo/`
- `│⠀⠀⠀└──⠀config.toml⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀#⠀WASM⠀optimization⠀flags`
- `└──⠀README.md`


### Generated `Cargo.toml` Pinning Matrix

The generator injects strict SDK versions to guarantee network compatibility:


- `[dependencies]`
- `soroban-sdk⠀=⠀"20.0.0"`
- `⠀`
- `[dev_dependencies]`
- `soroban-sdk⠀=⠀{⠀version⠀=⠀"20.0.0",⠀features⠀=⠀["testutils"]⠀}`
- `⠀`
- `[profile.release]`
- `opt-level⠀=⠀"z"`
- `overflow-checks⠀=⠀true`
- `debug⠀=⠀0`
- `strip⠀=⠀"symbols"`
- `debug-assertions⠀=⠀false`
- `panic⠀=⠀"abort"`
- `codegen-units⠀=⠀1`
- `lto⠀=⠀true`


---

## 6. CI/CD Integration Automation

By default, the scaffold injects a continuous integration pipeline utilizing `stellarpath-action`. This ensures that from day one, your repository enforces the strictest security standards.

The injected `stellarpath.yml` automatically scans PRs for `#17 panic` and `#18 unwrap` errors before allowing merges to `main`.

---

## 7. Advanced Makefile Deployment Automation

A key advantage of `stellar-scaffold` is the pre-configured `Makefile` it provisions. Instead of requiring developers to manually memorize `soroban-cli` arguments, the Makefile acts as an execution facade:


- `make⠀build⠀⠀⠀⠀⠀⠀⠀#⠀Compiles⠀WASM⠀to⠀target/wasm32-unknown-unknown/release/`
- `make⠀test⠀⠀⠀⠀⠀⠀⠀⠀#⠀Runs⠀Rust⠀integration⠀tests⠀with⠀testutils`
- `make⠀deploy⠀⠀⠀⠀⠀⠀#⠀Invokes⠀soroban⠀contract⠀deploy⠀on⠀Futurenet`
- `make⠀bindings⠀⠀⠀⠀#⠀Generates⠀TypeScript⠀bindings⠀for⠀frontend⠀clients`


This ensures that deploying a Soroban contract is as frictionless as standard web development workflows.

---

## 8. Contract Size Optimization Profiling

Because Soroban smart contracts have stringent size limits for mainnet deployment, `stellar-scaffold` optimizes the `.cargo/config.toml` from the start.

It automatically disables `debug-assertions`, forces `lto = true`, and sets `codegen-units = 1` so that the resulting `.wasm` binary is drastically smaller. The generated template is specifically tuned to compile securely without inflating the WASM footprint, preventing out-of-gas deployment errors right out of the box.

---
<div align="center">
  <sub>Part of the <b>STELLAR-PATH</b> Toolchain. Built for the Soroban ecosystem.</sub>
</div>
