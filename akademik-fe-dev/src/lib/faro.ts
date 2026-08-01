import { getWebInstrumentations, initializeFaro } from "@grafana/faro-web-sdk";

export function initFaro() {
  if (typeof window === "undefined") return;

  const faro = initializeFaro({
    url: process.env.NEXT_PUBLIC_FARO_URL,
    app: {
      name: "akademik-fe",
      version: "1.0.0",
      environment: process.env.NODE_ENV || "development",
    },
    instrumentations: [
      ...getWebInstrumentations({
        captureConsole: true,
        // captureConsoleDisabledLevels: [InternalLoggerLevel.DEBUG],
      }),
    ],
    beforeSend: (item) => {
      // Filter out sensitive data
      if (item.type === "log" && typeof item.payload === "object") {
        const payload = item.payload as any;
        if (payload.message && typeof payload.message === "string") {
          // Don't send logs containing sensitive keywords
          const sensitiveKeywords = ["password", "token", "key", "secret"];
          if (
            sensitiveKeywords.some((keyword) =>
              payload.message.toLowerCase().includes(keyword)
            )
          ) {
            return null; // Don't send this log
          }
        }
      }
      return item;
    },
  });

  return faro;
}
