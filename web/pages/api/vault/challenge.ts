import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  // Get API URL
  let domain = new URL(req.url).searchParams.get("domain");
  let username = req.headers.get("username");

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const resp = await fetch(
    process.env.API_URL + "/sonr-io/highway/vault/challenge/" + domain + "/" + username,
    requestOptions
  );
  const data = await resp.json();
  return new Response(
    JSON.stringify({
      session_id: data.session_id,
      creation_options: data.creation_options,
      domain: domain,
    }),
    {
      status: 200,
      headers: {
        "content-type": "application/json",
      },
    }
  );
}
