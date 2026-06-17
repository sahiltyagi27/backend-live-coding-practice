const crypto = require("node:crypto");
const { promisify } = require("node:util");

const pbkdf2 = promisify(crypto.pbkdf2);

async function hashPassword(i) {
  await pbkdf2(`password-${i}`, "salt", 100_000, 32, "sha256");
  return i;
}

async function main() {
  console.time("pbkdf2");
  const results = await Promise.all([1, 2, 3, 4].map(hashPassword));
  console.timeEnd("pbkdf2");
  console.log("completed tasks:", results);
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
