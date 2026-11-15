<div align="center">

# ⛓️ MAD BLOCKS

### Full blockchain implementation in Go — Bitcoin-style, with P2P consensus

[![Go](https://img.shields.io/badge/Go-1.20-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)

> *A from-scratch blockchain with Proof of Work, ECDSA wallets, and P2P consensus — no external blockchain libs.*

</div>

---

## 📖 Overview

**MAD BLOCKS** is a complete, from-scratch blockchain implementation in Go. It covers block creation, Proof of Work, signed transactions, wallet management, currency mining, P2P neighbor discovery, heartbeat sync, and longest-chain consensus. Each node runs an HTTP server exposing the chain, mining, transaction, and consensus endpoints.

## ✨ Features

- 🧱 **Block creation & chain verification** — SHA256-hashed blocks linked by previous hash
- ⛏️ **Proof of Work** — mining with configurable difficulty
- 💸 **Signed transactions** — ECDSA (P-256) signatures, validated on intake
- 💰 **Wallets** — Bitcoin-style address derivation (SHA256 → RIPEMD160 → base58, with checksum)
- ⚡ **Currency mining** — coinbase rewards to miner wallet
- 🌐 **P2P discovery** — subnet scan to find neighbor nodes
- 💓 **Heartbeat sync** — periodic neighbor re-discovery
- 🤝 **Consensus** — longest-chain resolution across neighbors
- 🔒 **Thread-safe** — mutex-guarded transaction pool and neighbor list
- 📡 **HTTP API** — get chain, mine, start mining, create transaction, balance, consensus

## 🏗 Architecture

```
mad-blocks/
├── block/          # Block, BlockChain, consensus, PoW, transactions
├── wallet/         # ECDSA keypair + Bitcoin-style address derivation
├── utils/          # hashing, neighbor discovery, HTTP helpers
├── hostServer/     # HTTP blockchain node (chain, mine, tx, consensus)
└── walletServer/   # wallet HTTP service
```

### HTTP endpoints (`hostServer`)

| Method | Path | Description |
|---|---|---|
| GET | `/` | Get full chain |
| GET | `/mine` | Mine a block |
| GET | `/mine/start` | Start continuous mining |
| GET | `/amount?address=` | Balance of address |
| GET/POST | `/transactions` | List / create transaction |
| PUT | `/consensus` | Resolve conflicts (longest chain) |

## 🛠 Tech stack

| Layer | Technology |
|---|---|
| Language | Go 1.20 |
| Crypto | ECDSA (P-256), SHA256, RIPEMD160, base58 |
| Networking | TCP (P2P discovery), HTTP (node API) |
| Concurrency | goroutines + sync.Mutex |

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/mad-blocks.git
cd mad-blocks

# run a blockchain node
go run hostServer/main.go

# run a wallet service (separate terminal)
go run walletServer/main.go
```

## 🗺 Roadmap

- [ ] Persistent chain storage (disk)
- [ ] Merkle tree for transaction verification
- [ ] Difficulty adjustment
- [ ] WebSocket-based P2P instead of polling
- [ ] Web dashboard for chain visualization

## 📄 License

MIT © [m3rashid](https://github.com/m3rashid)
