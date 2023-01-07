import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: req.body,
  };
  return await fetch(
    "https://api.sonr.network/sonr-io/highway/vault/register",
    requestOptions
  );
}
