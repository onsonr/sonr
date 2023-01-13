import type { NextRequest } from "next/server";

export default async function handler(req: NextRequest) {
  let body = req.body;
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: body,
  };
  const faucetUrl = process.env.FAUCET_URL;
  if (!faucetUrl) {
    return new Response(JSON.stringify({ error: "Faucet URL not set" }), {
      status: 500,
      headers: {
        "content-type": "application/json",
      },
    });
  }
  const resp = await fetch(faucetUrl, requestOptions);
  const data = await resp.json();
  console.log(data);
  return new Response(JSON.stringify(data), {
    status: 200,
    headers: {
      "content-type": "application/json",
    },
  });
}
