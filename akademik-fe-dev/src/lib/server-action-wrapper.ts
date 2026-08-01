import { headers } from "next/headers";
import { serverLogger, measureServerAction } from "./server-logger";

// Extract user info from headers (if available from auth)
const getUserInfoFromHeaders = () => {
  try {
    const headersList = headers();
    // You might get user info from JWT in cookies or authorization header
    // This is just an example - adjust based on your auth setup
    const authHeader = headersList.get("authorization");
    const userAgent = headersList.get("user-agent");
    const ip =
      headersList.get("x-forwarded-for") || headersList.get("x-real-ip");

    return {
      userAgent,
      ip,
      hasAuth: !!authHeader,
    };
  } catch (error) {
    return console.log(error);
  }
};

// Wrapper for server actions with automatic logging
export function withServerActionLogging<T extends any[], R>(
  actionName: string,
  action: (...args: T) => Promise<R>
) {
  return async (...args: T): Promise<R> => {
    const userInfo = getUserInfoFromHeaders();

    return measureServerAction(actionName, () => action(...args), {
      args: args.length,
      userInfo,
    });
  };
}

// Decorator for server actions (if you prefer decorator pattern)
export function logServerAction(actionName: string) {
  return function <T extends any[], R>(
    _target: any,
    propertyKey: string,
    descriptor: TypedPropertyDescriptor<(...args: T) => Promise<R>>
  ) {
    const originalMethod = descriptor.value!;

    descriptor.value = async function (...args: T): Promise<R> {
      const userInfo = getUserInfoFromHeaders();

      return measureServerAction(
        actionName || propertyKey,
        () => originalMethod.apply(this, args),
        {
          args: args.length,
          userInfo,
        }
      );
    };

    return descriptor;
  };
}

// Helper to log server action results
export const logActionResult = <T>(
  actionName: string,
  result: T,
  metadata?: Record<string, any>
) => {
  serverLogger.info(`Action result: ${actionName}`, {
    resultType: typeof result,
    hasResult: result !== null && result !== undefined,
    ...metadata,
  });
  return result;
};

// Helper to log server action errors
export const logActionError = (
  actionName: string,
  error: Error,
  metadata?: Record<string, any>
) => {
  serverLogger.error(`Action error: ${actionName}`, error, metadata);
  throw error; // Re-throw to maintain error flow
};
