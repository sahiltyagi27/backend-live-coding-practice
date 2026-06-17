const { Worker } = require("node:worker_threads");
const path = require("node:path");

function runWorker(n) {
  return new Promise((resolve, reject) => {
    const worker = new Worker(path.join(__dirname, "worker.js"), {
      workerData: n,
    });

    worker.on("message", resolve);
    worker.on("error", reject);
    worker.on("exit", (code) => {
      if (code !== 0) {
        reject(new Error(`worker exited with code ${code}`));
      }
    });
  });
}

async function main() {
  console.time("workers");
  const results = await Promise.all([35, 36, 37].map(runWorker));
  console.timeEnd("workers");
  console.log(results);
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
