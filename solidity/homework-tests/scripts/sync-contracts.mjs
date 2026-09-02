import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const here = path.dirname(fileURLToPath(import.meta.url));
const testsRoot = path.resolve(here, "..");
const contractsDir = path.join(testsRoot, "contracts");
const homework01Dir = path.resolve(testsRoot, "../homework01");
const homework02Dir = path.resolve(testsRoot, "../homework02");

fs.mkdirSync(contractsDir, { recursive: true });

for (const file of fs.readdirSync(contractsDir)) {
  if (file.endsWith(".sol")) {
    fs.unlinkSync(path.join(contractsDir, file));
  }
}

const sources = [
  ...fs.readdirSync(homework01Dir).filter((file) => file.endsWith(".sol")).map((file) => path.join(homework01Dir, file)),
  path.join(homework02Dir, "BeggingContract.sol"),
];

for (const source of sources) {
  fs.copyFileSync(source, path.join(contractsDir, path.basename(source)));
}

console.log(`synced ${sources.length} Solidity files into homework-tests/contracts`);
