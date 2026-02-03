<template>
  <div id="app" style="text-align: center; padding: 50px; font-family: sans-serif">
    <h1>Somnia Reactive Counter</h1>

    <div v-if="!account" style="margin-top: 20px">
      <button @click="connectWallet" style="padding: 10px 20px; cursor: pointer">
        Connect MetaMask
      </button>
    </div>

    <div v-else>
      <p>
        Connected: <b>{{ shortAddress }}</b>
      </p>

      <div style="font-size: 3rem; margin: 20px 0">Count: {{ count }}</div>

      <button
        @click="incrementCounter"
        :disabled="loading"
        style="padding: 15px 30px; font-size: 1.2rem; cursor: pointer"
      >
        {{ loading ? 'Processing...' : '➕ Increment' }}
      </button>

      <p v-if="txHash" style="margin-top: 20px; font-size: 0.8rem">
        Last Tx:
        <a :href="'https://shannon-explorer.somnia.network/tx/' + txHash" target="_blank"
          >View on Explorer</a
        >
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { ethers } from 'ethers'

// 1. DATA & STATE
const account = ref<string | null>(null)
const count = ref(0)
const loading = ref(false)
const txHash = ref('')

// Replace with your address from the Hardhat deployment!
const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS
const ABI = JSON.parse(import.meta.env.VITE_CONTRACT_ABI)

let contract: ethers.Contract | null = null

// 2. LOGIC
const shortAddress = computed(() => {
  return account.value ? `${account.value.slice(0, 6)}...${account.value.slice(-4)}` : ''
})

const connectWallet = async () => {
  if ((window as any).ethereum) {
    const provider = new ethers.BrowserProvider((window as any).ethereum)
    const accounts = await provider.send('eth_requestAccounts', [])
    account.value = accounts[0]

    const signer = await provider.getSigner()
    contract = new ethers.Contract(CONTRACT_ADDRESS, ABI, signer)

    fetchCount()
    listenToEvents()
  } else {
    alert('Please install MetaMask!')
  }
}

const fetchCount = async () => {
  if (!contract) return
  const currentCount = await contract.x?.()
  count.value = currentCount?.toString() || '0'
}

const listenToEvents = () => {
  // This is the "Reactive" part!
  // When Somnia emits the event, Vue updates the UI automatically.
  contract?.on('Increment', (newCount) => {
    console.log('Event detected on-chain!')
    count.value = newCount.toString()
  })
}

const incrementCounter = async () => {
  if (!contract) return
  try {
    loading.value = true
    const tx = await contract.inc?.()
    txHash.value = tx?.hash || ''
    await tx.wait() // Somnia is so fast this will feel nearly instant
    await fetchCount()
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  if ((window as any).ethereum?.selectedAddress) connectWallet()
})
</script>
