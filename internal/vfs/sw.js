importScripts(
  "https://cdn.jsdelivr.net/gh/golang/go@go1.18.4/misc/wasm/wasm_exec.js",
);

importScripts(
  "https://cdn.jsdelivr.net/gh/nlepage/go-wasm-http-server@v1.1.0/sw.js",
);

registerWasmHTTPListener("dwn.wasm");
