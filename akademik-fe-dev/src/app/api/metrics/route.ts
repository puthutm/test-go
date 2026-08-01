import client from "prom-client";

const collectDefaultMetrics = client.collectDefaultMetrics;
const Registry = client.Registry;
const register = new Registry();
collectDefaultMetrics({
  labels: { APP_NAME: "LMS" },
  register,
  prefix: "lms_",
});

export async function GET(req: Request) {
  const APIKEY = process.env.API_KEY;
  const url = new URL(req.url);

  const apiKeyParam = url.searchParams.get("api_key");

  if (!apiKeyParam) {
    return new Response("Access denied", {
      status: 403,
    });
  }

  if (apiKeyParam !== APIKEY) {
    return new Response("Invalid credentials", {
      status: 401,
    });
  }

  const metrics = await register.metrics();
  return new Response(metrics, {
    headers: {
      "Content-Type": register.contentType,
    },
  });
}
