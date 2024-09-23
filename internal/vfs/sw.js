importScripts(
  "https://cdn.jsdelivr.net/gh/golang/go@go1.18.4/misc/wasm/wasm_exec.js",
  "https://cdn.jsdelivr.net/gh/nlepage/go-wasm-http-server@v1.1.0/sw.js",
  "https://cdn.jsdelivr.net/npm/htmx.org@1.9.12/dist/htmx.min.js",
  "https://cdn.jsdelivr.net/npm/alpinejs@3.14.1/dist/cdn.min.js",
);

registerWasmHTTPListener("dwn.wasm");
