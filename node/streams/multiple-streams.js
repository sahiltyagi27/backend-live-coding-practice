const fs = require("fs");
const os = require("os");
const path = require("path");
const { Transform } = require("stream");
const { pipeline } = require("stream/promises");

function uppercaseTransform() {
  return new Transform({
    transform(chunk, encoding, callback) {
      callback(null, chunk.toString().toUpperCase());
    },
  });
}

async function streamFile(inputPath, outputPath) {
  await pipeline(
    fs.createReadStream(inputPath),
    uppercaseTransform(),
    fs.createWriteStream(outputPath),
  );
}

async function processMultipleStreams(files) {
  await Promise.all(
    files.map((file) => streamFile(file.input, file.output)),
  );
}

async function main() {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), "node-streams-"));

  const files = [1, 2, 3].map((id) => {
    const input = path.join(dir, `request-${id}.txt`);
    const output = path.join(dir, `request-${id}.out.txt`);
    fs.writeFileSync(input, `request ${id}\nlarge payload would stream here\n`);
    return { input, output };
  });

  console.time("streams");
  await processMultipleStreams(files);
  console.timeEnd("streams");

  for (const file of files) {
    const output = fs.readFileSync(file.output, "utf8").trim();
    console.log(path.basename(file.output), "=>", output.split("\n")[0]);
  }
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
