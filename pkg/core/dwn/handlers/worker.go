package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/core/dwn"
)

// generateRawServiceWorkerJS returns the service worker JavaScript as a string
func generateRawServiceWorkerJS(cfg *dwn.Environment) string {
	return fmt.Sprintf(`const CACHE_NAMES = {
    wasm: "wasm-cache-%s",
    static: "static-cache-%s",
    dynamic: "dynamic-cache-%s"
};

importScripts(
    %q,
    %q
);

// Initialize WASM HTTP listener with configured path
const wasmInstance = registerWasmHTTPListener(%q);

// MessageChannel port for WASM communication
let wasmPort;

// Request queue for offline operations
let requestQueue = new Map();

// Setup message channel handler
self.addEventListener('message', async (event) => {
    if (event.data.type === 'PORT_INITIALIZATION') {
        wasmPort = event.data.port;
        setupWasmCommunication();
    }
});

function setupWasmCommunication() {
    wasmPort.onmessage = async (event) => {
        const { type, data } = event.data;
        
        switch (type) {
            case 'WASM_REQUEST':
                handleWasmRequest(data);
                break;
            case 'SYNC_REQUEST':
                processSyncQueue();
                break;
        }
    };

    // Notify that WASM is ready
    wasmPort.postMessage({ type: 'WASM_READY' });
}

// Enhanced install event
self.addEventListener("install", (event) => {
    event.waitUntil(
        Promise.all([
            skipWaiting(),
            // Cache WASM binary and essential resources
            caches.open(CACHE_NAMES.wasm).then(cache => 
                cache.addAll([
                    %q,
                    %q
                ])
            )
        ])
    );
});

// Enhanced activate event
self.addEventListener("activate", (event) => {
    event.waitUntil(
        Promise.all([
            clients.claim(),
            // Clean up old caches
            caches.keys().then(keys => 
                Promise.all(
                    keys.map(key => {
                        if (!Object.values(CACHE_NAMES).includes(key)) {
                            return caches.delete(key);
                        }
                    })
                )
            )
        ])
    );
});

// Intercept fetch events
self.addEventListener('fetch', (event) => {
    const request = event.request;

    // Handle API requests differently from static resources
    if (request.url.includes('/api/')) {
        event.respondWith(handleApiRequest(request));
    } else {
        event.respondWith(handleStaticRequest(request));
    }
});

async function handleApiRequest(request) {
    try {
        // Try to make the request
        const response = await fetch(request.clone());
        
        // If successful, pass through WASM handler
        if (response.ok) {
            return await processWasmResponse(request, response);
        }
        
        // If offline or failed, queue the request
        await queueRequest(request);
        
        // Return cached response if available
        const cachedResponse = await caches.match(request);
        if (cachedResponse) {
            return cachedResponse;
        }
        
        // Return offline response
        return new Response(
            JSON.stringify({ error: 'Currently offline' }),
            { 
                status: 503,
                headers: { 'Content-Type': 'application/json' }
            }
        );
    } catch (error) {
        await queueRequest(request);
        return new Response(
            JSON.stringify({ error: 'Request failed' }),
            { 
                status: 500,
                headers: { 'Content-Type': 'application/json' }
            }
        );
    }
}

async function handleStaticRequest(request) {
    // Check cache first
    const cachedResponse = await caches.match(request);
    if (cachedResponse) {
        return cachedResponse;
    }

    try {
        const response = await fetch(request);
        
        // Cache successful responses
        if (response.ok) {
            const cache = await caches.open(CACHE_NAMES.static);
            cache.put(request, response.clone());
        }
        
        return response;
    } catch (error) {
        // Return offline page for navigation requests
        if (request.mode === 'navigate') {
            return caches.match('/offline.html');
        }
        throw error;
    }
}

async function processWasmResponse(request, response) {
    const responseClone = response.clone();
    
    try {
        const processedResponse = await wasmInstance.processResponse(responseClone);
        
        if (wasmPort) {
            wasmPort.postMessage({
                type: 'RESPONSE',
                requestId: request.headers.get('X-Wasm-Request-ID'),
                response: processedResponse
            });
        }
        
        return processedResponse;
    } catch (error) {
        console.error('WASM processing error:', error);
        return response;
    }
}

async function queueRequest(request) {
    const serializedRequest = await serializeRequest(request);
    requestQueue.set(request.url, serializedRequest);
    
    try {
        await self.registration.sync.register('wasm-sync');
    } catch (error) {
        console.error('Sync registration failed:', error);
    }
}

async function serializeRequest(request) {
    const headers = {};
    for (const [key, value] of request.headers.entries()) {
        headers[key] = value;
    }
    
    return {
        url: request.url,
        method: request.method,
        headers,
        body: await request.text(),
        timestamp: Date.now()
    };
}

// Handle background sync
self.addEventListener('sync', (event) => {
    if (event.tag === 'wasm-sync') {
        event.waitUntil(processSyncQueue());
    }
});

async function processSyncQueue() {
    const requests = Array.from(requestQueue.values());
    
    for (const serializedRequest of requests) {
        try {
            const response = await fetch(new Request(serializedRequest.url, {
                method: serializedRequest.method,
                headers: serializedRequest.headers,
                body: serializedRequest.body
            }));
            
            if (response.ok) {
                requestQueue.delete(serializedRequest.url);
                
                if (wasmPort) {
                    wasmPort.postMessage({
                        type: 'SYNC_COMPLETE',
                        url: serializedRequest.url
                    });
                }
            }
        } catch (error) {
            console.error('Sync failed for request:', error);
        }
    }
}

// Handle payment requests
self.addEventListener("canmakepayment", function (e) {
    e.respondWith(Promise.resolve(true));
});

// Handle periodic sync if available
self.addEventListener('periodicsync', (event) => {
    if (event.tag === 'wasm-sync') {
        event.waitUntil(processSyncQueue());
    }
});`,
		cfg.CacheVersion,
		cfg.CacheVersion,
		cfg.CacheVersion,
		cfg.WasmExecPath,
		cfg.HttpserverPath,
		cfg.WasmPath,
		cfg.WasmPath,
		cfg.WasmExecPath,
	)
}

// ServiceWorkerHandler is an Echo handler that serves the service worker
func ServiceWorkerHandler(cfg *dwn.Environment) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set appropriate headers for service worker
		c.Response().Header().Set("Content-Type", "application/javascript")
		c.Response().Header().Set("Service-Worker-Allowed", "/")

		// Generate and write the service worker JavaScript
		return c.String(http.StatusOK, generateRawServiceWorkerJS(cfg))
	}
}
