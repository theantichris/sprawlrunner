# Sprawlrunner

[![Go Reference](https://pkg.go.dev/badge/github.com/theantichris/sprawlrunner.svg)](https://pkg.go.dev/github.com/theantichris/sprawlrunner)
[![Build Status](https://github.com/theantichris/sprawlrunner/actions/workflows/go.yml/badge.svg)](https://github.com/theantichris/sprawlrunner/actions)
[![Build Status](https://github.com/theantichris/sprawlrunner/actions/workflows/markdown.yml/badge.svg)](https://github.com/theantichris/sprawlrunner/actions)
[![Go ReportCard](https://goreportcard.com/badge/theantichris/sprawrunner)](https://goreportcard.com/report/theantichris/sprawrunner)
![license](https://img.shields.io/badge/license-MIT-informational?style=flat)

A rogue-like ASCII game in a cyberpunk universe inspired by [ADOM](https://www.adom.de/home/index.html)
and [Shadowrun](https://store.catalystgamelabs.com/collections/shadowrun).

## Getting Started

### Prerequisites

- Go 1.23 or later ([Download Go](https://go.dev/dl/))
- Git

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/theantichris/sprawlrunner.git
   cd sprawlrunner
   ```

2. **Download dependencies**:

   ```bash
   go mod download
   ```

### Running the Game

Run the game directly with:

```bash
go run ./cmd/game
```

### Building

To build an executable:

```bash
go build -o sprawlrunner ./cmd/game
```

Then run the built executable:

- **Linux/macOS**: `./sprawlrunner`
- **Windows**: `sprawlrunner.exe`

## Controls

### Interface

| Action | Key |
| ------ | --- |
| Quit   | Q   |

### Movement

| Direction  | Vi | Numpad | Arrows/Nav |
| -----------| -- | ------ | ---------- |
| up left    | y  | 7      | Home       |
| up         | k  | 8      | ↑          |
| up right   | u  | 9      | PgUp       |
| left       | h  | 4      | ←          |
| right      | l  | 6      | →          |
| down left  | b  | 1      | End        |
| down       | j  | 2      | ↓          |
| down right | n  | 3      | PgDn       |
