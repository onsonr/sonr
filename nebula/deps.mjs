// deps.mjs
import { mkdir, writeFile } from "fs/promises";
import fetch from "node-fetch";
import path from "path";

async function fetchAndSave(url, outputPath) {
	try {
		const response = await fetch(url);
		if (!response.ok) {
			throw new Error(`Failed to fetch ${url}: ${response.statusText}`);
		}
		const data = await response.text();
		await writeFile(outputPath, data, "utf8");
		console.log(`Fetched and saved: ${outputPath}`);
	} catch (error) {
		console.error(`Error fetching ${url}:`, error);
	}
}

async function main() {
	// Ensure the assets directories exist
	await mkdir("./assets/js", { recursive: true });
	await mkdir("./assets/css", { recursive: true });

	// Fetch htmx.min.js
	await fetchAndSave(
		"https://cdnjs.cloudflare.com/ajax/libs/htmx/2.0.2/htmx.min.js",
		"./assets/js/htmx.min.js",
	);

	// Fetch alpine.min.js
	await fetchAndSave(
		"https://cdnjs.cloudflare.com/ajax/libs/alpinejs/3.14.1/cdn.min.js",
		"./assets/js/alpine.min.js",
	);

	// Fetch dexie.min.js
	await fetchAndSave(
		"https://cdnjs.cloudflare.com/ajax/libs/dexie/4.0.8/dexie.min.js",
		"./assets/js/dexie.min.js",
	);
}

main();
