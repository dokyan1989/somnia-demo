# Project Title

A full-stack application featuring a **Go** backend, a **Vue.js** frontend, and **Somnia** blockchain integration.

## 🚀 Getting Started

### Prerequisites

* **Go** (1.21+)
* **Node.js** (v22) & **npm**
* **jq** & **make**
* **abigen** (`make install-tools`)

---

## 🛠 Installation & Running

### 1. Smart Contract Bindings

Instead of manual commands, use the Makefile to sync your Go backend with your Solidity contracts:

```bash
# Install abigen if you haven't yet
make install-tools

# Generate counter.go from Hardhat artifacts
make build-contract

```

### 2. Backend (Go)

```bash
# Run the server
make run-backend

```

### 3. Frontend (Vue.js)

```bash
cd frontend
npm install
npm run dev

```

---

## ⛓ Blockchain Development

Smart contracts are managed via **Hardhat** on the **Somnia** network.

### Build & Deploy

```bash
# Compile contracts
npx hardhat build

# Deploy to Somnia
npx hardhat ignition deploy ignition/modules/ReactiveCounter.ts --network somnia

```

### ABI Management

Generate the raw ABI for the frontend:

```bash
npx tsx print-abi.ts

```

---

## ⚙️ Configuration

Ensure you have `.env` files in the appropriate directories. Refer to `.env.example` for the required keys.

---

### Important Note on the Makefile

The `Makefile` expects your directory structure to look like this:

```text
.
├── backend/
│   ├── main.go
│   └── counter.go (generated)
├── blockchain/
│   └── artifacts/
└── Makefile

```

## 🛠 Workflow

### One-Step Deploy & Sync

To deploy the contract to Somnia and automatically update the Go backend bindings, run:

```bash
./deploy.sh
```
