import { faro } from '@grafana/faro-web-sdk';

// Custom logging functions
export const faroLogger = {
  info: (message: string, meta?: Record<string, any>) => {
    if (faro.api) {
      faro.api.pushLog(['info', message, meta]);
    }
    console.info(message, meta);
  },

  warn: (message: string, meta?: Record<string, any>) => {
    if (faro.api) {
      faro.api.pushLog(['warn', message, meta]);
    }
    console.warn(message, meta);
  },

  error: (message: string, error?: Error, meta?: Record<string, any>) => {
    if (faro.api) {
      faro.api.pushLog(['error', message, { ...meta, error: error?.message, stack: error?.stack }]);
    }
    console.error(message, error, meta);
  },

  debug: (message: string, meta?: Record<string, any>) => {
    if (faro.api) {
      faro.api.pushLog(['debug', message, meta]);
    }
    console.debug(message, meta);
  },
};

// Custom event tracking
export const faroEvents = {
  // User actions
  userLogin: (userId: string, role?: string) => {
    if (faro.api) {
      faro.api.pushEvent('user_login', { userId, role: role || '' }, 'user');
    }
  },

  userLogout: (userId: string) => {
    if (faro.api) {
      faro.api.pushEvent('user_logout', { userId }, 'user');
    }
  },

  // Navigation
  pageView: (page: string, meta?: Record<string, any>) => {
    if (faro.api) {
      faro.api.pushEvent('page_view', { page, ...meta }, 'navigation');
    }
  },

  // API calls
  apiCall: (endpoint: string, method: string, status: number, duration?: number) => {
    if (faro.api) {
      faro.api.pushEvent('api_call', {
        endpoint,
        method,
        status: status.toString(),
        duration: duration?.toString() || ''
      }, 'api');
    }
  },

  // Errors
  apiError: (endpoint: string, method: string, error: string, status?: number) => {
    if (faro.api) {
      faro.api.pushEvent('api_error', {
        endpoint,
        method,
        error,
        status: status?.toString() || ''
      }, 'error');
    }
  },

  // Custom business events
  customEvent: (name: string, attributes: Record<string, any>, domain = 'custom') => {
    if (faro.api) {
      faro.api.pushEvent(name, attributes, domain);
    }
  },
};

// Performance tracking
export const faroPerformance = {
  measurePageLoad: () => {
    if (typeof window !== 'undefined' && faro.api) {
      window.addEventListener('load', () => {
        const navigation = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
        if (navigation) {
          faro.api.pushEvent('page_load_performance', {
            loadTime: (navigation.loadEventEnd - navigation.loadEventStart).toString(),
            domContentLoaded: (navigation.domContentLoadedEventEnd - navigation.domContentLoadedEventStart).toString(),
            firstContentfulPaint: (navigation.responseEnd - navigation.requestStart).toString(),
          }, 'performance');
        }
      });
    }
  },

  measureApiCall: async <T>(
    operation: () => Promise<T>,
    operationName: string
  ): Promise<T> => {
    const startTime = performance.now();
    try {
      const result = await operation();
      const duration = performance.now() - startTime;

      if (faro.api) {
        faro.api.pushEvent('api_performance', {
          operation: operationName,
          duration: duration.toString(),
          success: 'true',
        }, 'performance');
      }

      return result;
    } catch (error) {
      const duration = performance.now() - startTime;

      if (faro.api) {
        faro.api.pushEvent('api_performance', {
          operation: operationName,
          duration: duration.toString(),
          success: 'false',
          error: error instanceof Error ? error.message : 'Unknown error',
        }, 'performance');
      }

      throw error;
    }
  },
};

// User session tracking
export const faroSession = {
  setUser: (userId: string, attributes?: Record<string, any>) => {
    if (faro.api) {
      faro.api.setUser({ id: userId, ...attributes });
    }
  },

  clearUser: () => {
    if (faro.api) {
      faro.api.setUser({});
    }
  },

  setSessionAttribute: (key: string, value: any) => {
    if (faro.api && 'setSessionAttribute' in faro.api) {
      (faro.api as any).setSessionAttribute(key, value);
    }
  },
};