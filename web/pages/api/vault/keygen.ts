import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
  };
  const resp = await fetch(
    "http://localhost:1317/sonr-io/highway/vault/keygen",
    requestOptions
  );
  const data = await resp.json();
  return new Response(JSON.stringify(data), {
    status: 200,
    headers: {
      "content-type": "application/json",
    },
  });
}
