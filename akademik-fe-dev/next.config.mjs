/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "example.com",
        port: "",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "svc-sso.dev-unsia.id",
        port: "",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "staging-svc-sso.dev-unsia.id",
        port: "",
        pathname: "/**",
      },
      {
        protocol: "https",
        hostname: "svc-sso.dev-unsia.idsso-be",
        port: "",
        pathname: "/**",
      },
    ],
  },
  async rewrites() {
    return [
      {
        source: "/metrics",
        destination: "/api/metrics",
      },
    ];
  },
  output: "standalone",
};

export default nextConfig;
