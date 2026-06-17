const { parentPort, workerData } = require("node:worker_threads");

function fib(n) {
  if (n <= 1) return n;
  return fib(n - 1) + fib(n - 2);
}

parentPort.postMessage({
  input: workerData,
  result: fib(workerData),
});
