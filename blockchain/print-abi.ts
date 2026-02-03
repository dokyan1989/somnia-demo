import { ethers } from "ethers";
import { createRequire } from "module"; // <--- 1. Import this

// 2. Create a 'require' function for this file
const require = createRequire(import.meta.url);

// 3. Now you can use require() just like before!
const artifact = require("./artifacts/contracts/ReactiveCounter.sol/ReactiveCounter.json");

async function main() {
  const iface = new ethers.Interface(artifact.abi);
  const humanReadable = iface.format(true);

  console.log("---------------------------------------------------");
  if (Array.isArray(humanReadable)) {
    humanReadable.forEach((line) => console.log(`"${line}",`));
  } else {
    console.log(humanReadable);
  }
  console.log("---------------------------------------------------");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});