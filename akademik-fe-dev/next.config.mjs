import { withSentryConfig } from "@sentry/nextjs";
/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "example.com",
        port: "", // Biarkan kosong jika tidak menggunakan port khusus
        pathname: "/**", // Izinkan semua path di domain ini
      },
      {
        protocol: "https",
        hostname: "svc-sso.dev-unsia.id",
        port: "", // Biarkan kosong jika tidak menggunakan port khusus
        pathname: "/**", // Izinkan semua path di domain ini
      },
      {
        protocol: "https",
        hostname: "staging-svc-sso.dev-unsia.id",
        port: "", // Biarkan kosong jika tidak menggunakan port khusus
        pathname: "/**", // Izinkan semua path di domain ini
      },
      {
        protocol: "https",
        hostname: "svc-sso.dev-unsia.idsso-be",
        port: "", // Biarkan kosong jika tidak menggunakan port khusus
        pathname: "/**", // Izinkan semua path di domain ini
      },
      //
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

export default withSentryConfig(nextConfig, {
  // For all available options, see:
  // https://www.npmjs.com/package/@sentry/webpack-plugin#options

  org: "sentry-unsia",
  project: "akademik-fe",
  sentryUrl: "https://sentry.unsia.ac.id/",

  // Only print logs for uploading source maps in CI
  silent: !process.env.CI,

  // For all available options, see:
  // https://docs.sentry.io/platforms/javascript/guides/nextjs/manual-setup/

  // Upload a larger set of source maps for prettier stack traces (increases build time)
  widenClientFileUpload: true,

  // Route browser requests to Sentry through a Next.js rewrite to circumvent ad-blockers.
  // This can increase your server load as well as your hosting bill.
  // Note: Check that the configured route will not match with your Next.js middleware, otherwise reporting of client-
  // side errors will fail.
  tunnelRoute: "/monitoring",

  // Automatically tree-shake Sentry logger statements to reduce bundle size
  disableLogger: true,

  // Enables automatic instrumentation of Vercel Cron Monitors. (Does not yet work with App Router route handlers.)
  // See the following for more information:
  // https://docs.sentry.io/product/crons/
  // https://vercel.com/docs/cron-jobs
  automaticVercelMonitors: true,

  authToken: process.env.SENTRY_AUTH_TOKEN,
});
