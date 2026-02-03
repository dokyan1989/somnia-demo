#!/bin/bash

# Exit on any error
set -e

echo "🚀 Starting Deployment to Somnia..."

# 1. Build and Deploy
npx hardhat build
npx hardhat ignition deploy ignition/modules/ReactiveCounter.ts --network somnia

echo "✅ Deployment Successful."

# 2. Sync Go Backend
echo "🔄 Syncing Go bindings..."
make build-contract

echo "✨ All systems synced and ready!"