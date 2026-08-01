// Server-side logging utilities that can send data to external services
// This can be extended to send to Grafana, Sentry, or other logging services

interface LogEntry {
  level: 'info' | 'warn' | 'error' | 'debug';
  message: string;
  metadata?: Record<string, any>;
  timestamp: string;
  environment: string;
  service: string;
}

interface ApiCallLog {
  url: string;
  method: string;
  status?: number;
  duration: number;
  error?: string;
  requestId?: string;
  userId?: string;
  timestamp: string;
}

// Send logs to Grafana Loki (same as client-side Faro)
const sendToExternalService = async (data: any) => {
  // Always console log for local debugging
  if (process.env.NODE_ENV === 'development') {
    console.log('[SERVER-LOG]', JSON.stringify(data, null, 2));
  }

  // Send to Grafana Loki if URL is configured
  const lokiUrl = process.env.NEXT_PUBLIC_FARO_URL;
  if (lokiUrl) {
    try {
      // Format data for Loki
      const lokiPayload = {
        streams: [
          {
            stream: {
              job: 'akademik-fe-server',
              level: data.level || 'info',
              service: data.service || 'akademik-fe',
              environment: data.environment || 'development',
              source: 'server',
            },
            values: [
              [
                (Date.now() * 1000000).toString(), // Loki expects nanoseconds
                JSON.stringify({
                  message: data.message || JSON.stringify(data),
                  ...data,
                })
              ]
            ]
          }
        ]
      };

      await fetch(lokiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(lokiPayload),
      });
    } catch (error) {
      // Don't break the app if logging fails
      console.error('Failed to send server log to Loki:', error);
    }
  }
};

export const serverLogger = {
  info: (message: string, metadata?: Record<string, any>) => {
    const logEntry: LogEntry = {
      level: 'info',
      message,
      metadata,
      timestamp: new Date().toISOString(),
      environment: process.env.NODE_ENV || 'development',
      service: 'akademik-fe',
    };
    sendToExternalService(logEntry);
  },

  warn: (message: string, metadata?: Record<string, any>) => {
    const logEntry: LogEntry = {
      level: 'warn',
      message,
      metadata,
      timestamp: new Date().toISOString(),
      environment: process.env.NODE_ENV || 'development',
      service: 'akademik-fe',
    };
    sendToExternalService(logEntry);
  },

  error: (message: string, error?: Error, metadata?: Record<string, any>) => {
    const logEntry: LogEntry = {
      level: 'error',
      message,
      metadata: {
        ...metadata,
        error: error?.message,
        stack: error?.stack,
        name: error?.name,
      },
      timestamp: new Date().toISOString(),
      environment: process.env.NODE_ENV || 'development',
      service: 'akademik-fe',
    };
    sendToExternalService(logEntry);
  },

  debug: (message: string, metadata?: Record<string, any>) => {
    if (process.env.NODE_ENV === 'development') {
      const logEntry: LogEntry = {
        level: 'debug',
        message,
        metadata,
        timestamp: new Date().toISOString(),
        environment: process.env.NODE_ENV || 'development',
        service: 'akademik-fe',
      };
      sendToExternalService(logEntry);
    }
  },

  // API call logging
  apiCall: (data: Omit<ApiCallLog, 'timestamp'>) => {
    const apiLog: ApiCallLog = {
      ...data,
      timestamp: new Date().toISOString(),
    };
    sendToExternalService({
      type: 'api_call',
      ...apiLog,
    });
  },
};

// Performance tracking for server actions
export const measureServerAction = async <T>(
  actionName: string,
  action: () => Promise<T>,
  metadata?: Record<string, any>
): Promise<T> => {
  const startTime = Date.now();
  const requestId = Math.random().toString(36).substring(7);

  serverLogger.info(`Starting server action: ${actionName}`, {
    requestId,
    ...metadata,
  });

  try {
    const result = await action();
    const duration = Date.now() - startTime;

    serverLogger.info(`Server action completed: ${actionName}`, {
      requestId,
      duration,
      success: true,
      ...metadata,
    });

    return result;
  } catch (error) {
    const duration = Date.now() - startTime;

    serverLogger.error(`Server action failed: ${actionName}`, error as Error, {
      requestId,
      duration,
      success: false,
      ...metadata,
    });

    throw error;
  }
};