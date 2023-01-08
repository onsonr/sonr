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
  const resp = await fetch(
    "http://localhost:1317/sonr-io/highway/vault/register",
    requestOptions
  );
  return resp;
}
