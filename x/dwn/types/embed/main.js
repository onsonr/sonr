// MessageChannel for WASM communication
let wasmChannel;
let wasmPort;

async function initWasmChannel() {
  wasmChannel = new MessageChannel();
  wasmPort = wasmChannel.port1;

  // Setup message handling from WASM
  wasmPort.onmessage = (event) => {
    const { type, data } = event.data;
    switch (type) {
      case "WASM_READY":
        console.log("WASM is ready");
        document.dispatchEvent(new CustomEvent("wasm-ready"));
        break;
      case "RESPONSE":
        handleWasmResponse(data);
        break;
      case "SYNC_COMPLETE":
        handleSyncComplete(data);
        break;
    }
  };
}

// Initialize WebAssembly and Service Worker
async function init() {
  try {
    // Register service worker
    if ("serviceWorker" in navigator) {
      const registration = await navigator.serviceWorker.register("./sw.js");
      console.log("ServiceWorker registered");

      // Wait for the service worker to be ready
      await navigator.serviceWorker.ready;

      // Initialize MessageChannel
      await initWasmChannel();

      // Send the MessageChannel port to the service worker
      navigator.serviceWorker.controller.postMessage(
        {
          type: "PORT_INITIALIZATION",
          port: wasmChannel.port2,
        },
        [wasmChannel.port2],
      );

      // Register for periodic sync if available
      if ("periodicSync" in registration) {
        try {
          await registration.periodicSync.register("wasm-sync", {
            minInterval: 24 * 60 * 60 * 1000, // 24 hours
          });
        } catch (error) {
          console.log("Periodic sync could not be registered:", error);
        }
      }
    }

    // Initialize HTMX with custom config
    htmx.config.withCredentials = true;
    htmx.config.wsReconnectDelay = "full-jitter";

    // Override HTMX's internal request handling
    htmx.config.beforeRequest = function (config) {
      // Add request ID for tracking
      const requestId = "req_" + Date.now();
      config.headers["X-Wasm-Request-ID"] = requestId;

      // If offline, handle through service worker
      if (!navigator.onLine) {
        return false; // Let service worker handle it
      }
      return true;
    };

    // Handle HTMX after request
    htmx.config.afterRequest = function (config) {
      // Additional processing after request if needed
    };

    // Handle HTMX errors
    htmx.config.errorHandler = function (error) {
      console.error("HTMX Error:", error);
    };
  } catch (error) {
    console.error("Initialization failed:", error);
  }
}

function handleWasmResponse(data) {
  const { requestId, response } = data;
  // Process the WASM response
  // This might update the UI or trigger HTMX swaps
  const targetElement = document.querySelector(
    `[data-request-id="${requestId}"]`,
  );
  if (targetElement) {
    htmx.process(targetElement);
  }
}

function handleSyncComplete(data) {
  const { url } = data;
  // Handle successful sync
  // Maybe refresh the relevant part of the UI
  htmx.trigger("body", "sync:complete", { url });
}

// Handle offline status changes
window.addEventListener("online", () => {
  document.body.classList.remove("offline");
  // Trigger sync when back online
  if (wasmPort) {
    wasmPort.postMessage({ type: "SYNC_REQUEST" });
  }
});

window.addEventListener("offline", () => {
  document.body.classList.add("offline");
});

// Custom event handlers for HTMX
document.addEventListener("htmx:beforeRequest", (event) => {
  const { elt, xhr } = event.detail;
  // Add request tracking
  const requestId = xhr.headers["X-Wasm-Request-ID"];
  elt.setAttribute("data-request-id", requestId);
});

document.addEventListener("htmx:afterRequest", (event) => {
  const { elt, successful } = event.detail;
  if (successful) {
    elt.removeAttribute("data-request-id");
  }
});

// Initialize everything when the page loads
document.addEventListener("DOMContentLoaded", init);

// Export functions that might be needed by WASM
window.wasmBridge = {
  triggerUIUpdate: function (selector, content) {
    const target = document.querySelector(selector);
    if (target) {
      htmx.process(
        htmx.parse(content).forEach((node) => target.appendChild(node)),
      );
    }
  },

  showNotification: function (message, type = "info") {
    // Implement notification system
    console.log(`${type}: ${message}`);
  },
};
