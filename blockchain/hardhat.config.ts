import { defineConfig } from "hardhat/config";
import hardhatIgnitionViemPlugin from "@nomicfoundation/hardhat-ignition-viem";
import "dotenv/config"; // Loads variables from .env into process.env

// Access variables from process.env
const SOMNIA_RPC_URL = process.env.SOMNIA_RPC_URL || "";
const METAMASK_PRIVATE_KEY = process.env.METAMASK_PRIVATE_KEY || "";

export default defineConfig({
  plugins: [hardhatIgnitionViemPlugin],
  solidity: {
    version: "0.8.28",
  },
  networks: {
    somnia: {
      type: "http",
      url: SOMNIA_RPC_URL,
      chainId: 50312,
      accounts: [METAMASK_PRIVATE_KEY], // Your MetaMask private key
    },
  },
});
