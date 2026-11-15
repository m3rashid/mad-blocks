# MAD BLOCKS

A full blockchain implementation in Go, Bitcoin-style, with P2P consensus.

## Features

- Block creation & chain verification, Proof of Work
- Transactions & validation, ECDSA signing (SHA256 + RIPEMD160, base58 addresses)
- Currency mining & wallet management
- P2P neighbor discovery via subnet scan, heartbeat sync, conflict resolution
- HTTP API: get chain, mine, start mining, create transaction, balance, consensus
- Mutex-guarded atomic transactions; modular (block / wallet / hostServer / walletServer)

## Architecture

- `block/` — Block, BlockChain, consensus, PoW, transactions
- `wallet/` — ECDSA keypair generation, Bitcoin-style address derivation
- `utils/` — hashing, neighbors discovery, HTTP helpers
- `hostServer/` — HTTP blockchain node
- `walletServer/` — wallet HTTP service

## Tech

Go · SHA256 · ECDSA · P2P (TCP) · HTTP

## Quick start

```bash
git clone https://github.com/m3rashid/mad-blocks.git
cd mad-blocks
go run hostServer/main.go   # or: go run walletServer/main.go
```

## License

MIT
