import type { NextRequest } from "next/server";
import { RegisterRequest } from "@buf/sonr-hq_sonr.grpc_web/highway/vault/v1/api_pb";
import axios from "axios";
export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  // Get API URL
  let apiUrl = "https://api.sonr.network";
  if (process && process.env.NODE_ENV === "development") {
    apiUrl = "http://localhost:1317";
  }

  let body = req.body;
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: body,
  };
  const resp = await fetch(
    process.env.API_URL + "/sonr-io/sonr/vault/new-wallet",
    requestOptions
  );
  const data = await resp.json();
  console.log(data);
  return new Response(JSON.stringify(data), {
    status: 200,
    headers: {
      "content-type": "application/json",
    },
  });
}
