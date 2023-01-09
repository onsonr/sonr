import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  // Get API URL
  let apiUrl = "https://api.sonr.network";
  if (process && process.env.NODE_ENV === "development") {
    apiUrl = "http://localhost:1317";
  }
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: req.body,
  };
  const resp = await fetch(
    apiUrl + "/sonr-io/highway/vault/register",
    requestOptions
  );
  const data = await resp.json();
  console.log(data);
  return resp;
}
