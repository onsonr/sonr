importScripts(
  "https://cdn.jsdelivr.net/gh/golang/go@go1.22.5/misc/wasm/wasm_exec.js",
  "https://cdn.jsdelivr.net/gh/nlepage/go-wasm-http-server@v1.1.0/sw.js",
  "https://storage.googleapis.com/workbox-cdn/releases/3.6.3/workbox-sw.js",
);

registerWasmHTTPListener("app.wasm");

// Skip installed stage and jump to activating stage
addEventListener("install", (event) => {
  event.waitUntil(skipWaiting());
});

// Start controlling clients as soon as the SW is activated
addEventListener("activate", (event) => {
  event.waitUntil(clients.claim());
});
