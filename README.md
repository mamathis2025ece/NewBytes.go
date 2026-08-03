# NewBytes.go
Port of the original bytes.js library from JavaScript to Go

An idiomatic Go port of the original JavaScript library **bytes.js** by visionmedia, created as part of the **Code Resurrection 2026 Hackathon**.

## Original Repository

https://github.com/visionmedia/bytes.js

## Overview

This project ports the core functionality of the original JavaScript implementation into Go while preserving its behavior.

The implementation provides functionality to:

- Parse human-readable byte strings into numeric byte values.
- Format numeric byte values into human-readable strings.
- Support units:
  - B
  - KB
  - MB
  - GB
  - TB
  - PB

## Project Structure

```
.
├── src/
│   ├── index.go
│   ├── index_test.go
│   └── go.mod
├── tests/
│   ├── original/
│   └── port/
├── fuzz/
├── bench/
├── README.md
├── DECISIONS.md
├── Dockerfile
└── port-mortem.toml
```

## Requirements

- Go 1.25 or later

## Build

```bash
cd src
go mod tidy
```

## Run Tests

```bash
cd src
go test -v
```

## Example

```go
size, _ := Parse("5MB")
fmt.Println(size)
// 5242880

fmt.Println(Format(1536, nil))
// 1.5KB
```

## Notes

This project was created for educational purposes as part of the Code Resurrection 2026 Hackathon and is intended to preserve the observable behavior of the original JavaScript implementation using idiomatic Go.
