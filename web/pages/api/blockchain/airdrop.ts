import type { NextRequest } from "next/server";

export default async function handler(req: NextRequest) {
  // Get API URL
  let apiUrl = "https://faucet.sonr.network";
  if (process && process.env.NODE_ENV === "development") {
    apiUrl = "http://localhost:4500/";
  }

  let body = req.body;
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: body,
  };
  const resp = await fetch(apiUrl, requestOptions);
  const data = await resp.json();
  console.log(data);
  return new Response(JSON.stringify(data), {
    status: 200,
    headers: {
      "content-type": "application/json",
    },
  });
}
