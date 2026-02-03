export {}; // Make this a module

declare global {
  interface Window {
    ethereum: import('ethers').Eip1193Provider; // Or strictly: import('ethers').Eip1193Provider
  }
}