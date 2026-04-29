# Distributed File System

A peer-to-peer distributed file storage system built in Go. This project implements a content-addressable storage (CAS) layer with a modular TCP-based networking foundation designed for distributed file transfer across nodes.

## Architecture

```
FileServer        -- High-level API and orchestration
Store             -- Content-addressable file persistence (SHA1-based)
Transport (p2p)   -- Pluggable network layer (TCP)
```

**Key design decisions:**
- Interface-driven architecture for swappable transports (TCP, UDP, WebSockets)
- Content-addressable storage using SHA1 hashing with hierarchical directory structure
- Pluggable encoding (GOB, raw bytes) and handshake protocols
- Functional options pattern for configuration

## Project Structure

```
.
├── main.go                 # Application entry point
├── server.go               # FileServer integrating storage and networking
├── store.go                # CAS file storage implementation
├── store_test.go           # Storage tests
├── Makefile                # Build automation
├── p2p/
│   ├── transport.go        # Transport and Peer interfaces
│   ├── tcp_transport.go    # TCP transport implementation
│   ├── tcp_transport_test.go
│   ├── encoding.go         # Message encoding/decoding
│   ├── handshake.go        # Peer handshake protocol
│   └── message.go          # RPC message definitions
└── bin/                    # Compiled binaries
```

## Getting Started

### Prerequisites

- Go 1.25+

### Build

```bash
make build
```

### Run

```bash
make run
```

The server listens on `:3000` by default.

### Test

```bash
make test
```

## How It Works

**Storage:** Files are stored using content-addressable paths derived from SHA1 hashes. Keys are hashed and split into 5-character directory segments, creating a balanced directory tree on disk.

**Networking:** The P2P transport layer handles TCP connections between nodes. Incoming messages are consumed asynchronously via Go channels. The transport supports custom handshake functions and message decoders.

**FileServer:** Combines the storage and transport layers, providing the interface for distributed file operations across peers.

## Dependencies

- [testify](https://github.com/stretchr/testify) -- testing assertions

## License

This project is open source.
