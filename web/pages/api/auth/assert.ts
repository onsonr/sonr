import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  // Get API URL
  let domain = new URL(req.url).searchParams.get("domain");
  let apiUrl = "https://api.sonr.network";
  if (process && process.env.NODE_ENV === "development") {
    apiUrl = "http://localhost:1317";
  }
  let username = req.headers.get("username");

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const resp = await fetch(
    apiUrl + "/sonr/protocol/auth/assertion/" + domain + "/" + username,
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
