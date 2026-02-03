import * as fs from 'fs';
import * as path from 'path';
import { fileURLToPath } from 'url'; // 👈 Add this
import { ethers } from 'ethers';

// 1. Manually define __dirname for ES Modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Now your existing code will work:
const ARTIFACT_PATH = path.resolve(__dirname, '../blockchain/artifacts/contracts/ReactiveCounter.sol/ReactiveCounter.json');
const ENV_PATH = path.resolve(__dirname, './.env');

const syncHumanReadableAbi = (): void => {
  try {
    if (!fs.existsSync(ARTIFACT_PATH)) {
      throw new Error(`Artifact not found at ${ARTIFACT_PATH}`);
    }

    // 1. Read JSON Artifact
    const artifact = JSON.parse(fs.readFileSync(ARTIFACT_PATH, 'utf8'));
    
    // 2. Convert to Human-Readable Format (string[])
    const iface = new ethers.Interface(artifact.abi);
    const humanReadable = iface.format(true); // This returns string[]

    // 3. Prepare .env string
    // We stringify the array so it can be parsed back correctly in Vue
    const abiValue = JSON.stringify(humanReadable);
    const envKey = 'VITE_CONTRACT_ABI';
    const newEnvLine = `${envKey}='${abiValue}'`;

    // 4. Update .env content
    const envContent = fs.existsSync(ENV_PATH) ? fs.readFileSync(ENV_PATH, 'utf8') : '';
    const envRegex = new RegExp(`^${envKey}=.*$`, 'm');

    let updatedContent: string;
    if (envRegex.test(envContent)) {
      updatedContent = envContent.replace(envRegex, newEnvLine);
    } else {
      updatedContent = envContent.trim() + `\n${newEnvLine}\n`;
    }

    fs.writeFileSync(ENV_PATH, updatedContent.trim() + '\n');
    
    console.log('✅ [Success]: Human-Readable ABI synced!');
    console.log('📄 Format:', humanReadable);
  } catch (error) {
    console.error('❌ [Error]:', error instanceof Error ? error.message : error);
    process.exit(1);
  }
};

syncHumanReadableAbi();