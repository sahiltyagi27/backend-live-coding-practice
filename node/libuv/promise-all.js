const fs = require("node:fs/promises");

async function main() {
  const files = [__filename, __filename, __filename];

  console.time("concurrent-read");
  const results = await Promise.all(files.map((file) => fs.readFile(file, "utf8")));
  console.timeEnd("concurrent-read");

  console.log("files read:", results.length);
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
