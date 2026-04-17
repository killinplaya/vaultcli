# CLI Password Manager

A command-line password manager written in Go with encryption, JSON-based storage, and clean architecture.

## About

Pet project built to practice Go and consolidate knowledge of clean architecture principles. Passwords are stored locally in an encrypted JSON file.

## Project Structure

```
.
├── main.go
├── go.mod
└── internal
    ├── app/         # Application logic
    ├── entity/      # Core domain entities with validation
    ├── storage/     # Storage interface & implementation
    └── utils/       # ID generation, time helpers, slice utilities
```

## Architecture

The project follows a layered architecture with separation of concerns:

- **Entity** — core data structures and validation
- **Storage** — abstracted behind an interface, decoupled from business logic
- **App** — orchestrates interaction between layers
- **Utils** — shared helpers (ID generation, time, slice operations)

## Tech Stack

- **Language:** Go
- **Storage:** JSON
- **Security:** Encryption at rest

## Status

🚧 Work in progress