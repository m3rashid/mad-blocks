<div align="center">

# ⛓️ MAD BLOCKS

**Full blockchain implementation in Go — Bitcoin-style, with P2P consensus.**

[![Go](https://img.shields.io/badge/Go-1.20-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

</div>

---

## ✨ Features

- 🧱 Block creation & chain verification, Proof of Work
- 💸 Transactions & validation, ECDSA signing (SHA256 + RIPEMD160, base58 addresses)
- ⛏️ Currency mining & wallet management
- 🌐 P2P neighbor discovery via subnet scan, heartbeat sync, conflict resolution
- 📡 HTTP API: get chain, mine, start mining, create transaction, balance, consensus
- 🔒 Mutex-guarded atomic transactions

## 🏗 Architecture

| Package | Purpose |
|---|---|
| `block/` | Block, BlockChain, consensus, PoW, transactions |
| `wallet/` | ECDSA keypair + Bitcoin-style address derivation |
| `utils/` | hashing, neighbor discovery, HTTP helpers |
| `hostServer/` | HTTP blockchain node |
| `walletServer/` | wallet HTTP service |

## 🛠 Tech

`Go` · `SHA256` · `ECDSA` · `P2P (TCP)` · `HTTP`

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/mad-blocks.git
cd mad-blocks
go run hostServer/main.go   # or: go run walletServer/main.go
```

## 📄 License

MIT
