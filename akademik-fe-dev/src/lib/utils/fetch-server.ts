"use server";

import { getServerSession } from "next-auth";

import authOptions from "@/config/next-auth";
// import { serverLogger } from "../server-logger";

interface FetchApiOptions extends RequestInit {
  cacheType?: RequestCache;
}

export const fetchApi = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const baseUrl = `${process.env.NEXT_PUBLIC_API_BASE_URL}/api`;
  const url = `${baseUrl}${endpoint}`;
  // const startTime = Date.now();
  // const method = (options.method || "GET").toUpperCase();
  // const requestId = Math.random().toString(36).substring(7);

  // Log request start
  // serverLogger.info(`Server API call: ${method} ${endpoint}`, {
  //   requestId,
  //   method,
  //   endpoint,
  //   url,
  //   hasAuth: !!token,
  //   userId: session?.user?.id,
  // });

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    // const duration = Date.now() - startTime;

    // Log API call metrics
    // serverLogger.apiCall({
    //   url: endpoint,
    //   method,
    //   status: response.status,
    //   duration,
    //   requestId,
    //   userId: session?.user?.id || undefined,
    // });

    if (response.status === 401) {
      // serverLogger.warn(`Unauthorized API call: ${method} ${endpoint}`, {
      //   requestId,
      //   status: 401,
      //   duration,
      // });
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();

    // serverLogger.info(`Server API call completed: ${method} ${endpoint}`, {
    //   requestId,
    //   status: response.status,
    //   duration,
    //   success: response.ok,
    // });

    return result;
  } catch (error: any) {
    // const duration = Date.now() - startTime;

    // serverLogger.error(`Server API call failed: ${method} ${endpoint}`, error, {
    //   requestId,
    //   duration,
    //   url,
    //   userId: session?.user?.id || undefined,
    // });

    console.error("Fetch error:", error);
    throw new Error(error?.message || "Something went wrong");
  }
};

export const fetchApiSso = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const baseUrl = `${process.env.NEXT_PUBLIC_API_SSO_URL}/api`;
  const url = `${baseUrl}${endpoint}`;
  // const startTime = Date.now();
  // const method = (options.method || "GET").toUpperCase();
  // const requestId = Math.random().toString(36).substring(7);

  // serverLogger.info(`Server SSO API call: ${method} ${endpoint}`, {
  //   requestId,
  //   method,
  //   endpoint,
  //   url,
  //   hasAuth: !!token,
  //   userId: session?.user?.id,
  // });

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    // const duration = Date.now() - startTime;

    // serverLogger.apiCall({
    //   url: endpoint,
    //   method,
    //   status: response.status,
    //   duration,
    //   requestId,
    //   userId: session?.user?.id || undefined,
    // });

    if (response.status === 401) {
      // serverLogger.warn(`Unauthorized SSO API call: ${method} ${endpoint}`, {
      //   requestId,
      //   status: 401,
      //   duration,
      // });
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();

    // serverLogger.info(`Server SSO API call completed: ${method} ${endpoint}`, {
    //   requestId,
    //   status: response.status,
    //   duration,
    //   success: response.ok,
    // });

    return result;
  } catch (error: any) {
    // const duration = Date.now() - startTime;

    // serverLogger.error(`Server SSO API call failed: ${method} ${endpoint}`, error, {
    //   requestId,
    //   duration,
    //   url,
    //   userId: session?.user?.id || undefined,
    // });

    console.error("Fetch error:", error);
    throw new Error(error?.message || "Something went wrong");
  }
};

export const fetchApiDatareferensi = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const baseUrl = `${process.env.NEXT_PUBLIC_API_DATA_REFERENSI_URL}/api`;
  const url = `${baseUrl}${endpoint}`;
  // const startTime = Date.now();
  // const method = (options.method || "GET").toUpperCase();
  // const requestId = Math.random().toString(36).substring(7);

  // serverLogger.info(`Server DataRef API call: ${method} ${endpoint}`, {
  //   requestId,
  //   method,
  //   endpoint,
  //   url,
  //   hasAuth: !!token,
  //   userId: session?.user?.id,
  // });

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    // const duration = Date.now() - startTime;

    // serverLogger.apiCall({
    //   url: endpoint,
    //   method,
    //   status: response.status,
    //   duration,
    //   requestId,
    //   userId: session?.user?.id || undefined,
    // });

    if (response.status === 401) {
      // serverLogger.warn(`Unauthorized DataRef API call: ${method} ${endpoint}`, {
      //   requestId,
      //   status: 401,
      //   duration,
      // });
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();

    // serverLogger.info(`Server DataRef API call completed: ${method} ${endpoint}`, {
    //   requestId,
    //   status: response.status,
    //   duration,
    //   success: response.ok,
    // });

    return result;
  } catch (error: any) {
    // const duration = Date.now() - startTime;

    // serverLogger.error(`Server DataRef API call failed: ${method} ${endpoint}`, error, {
    //   requestId,
    //   duration,
    //   url,
    //   userId: session?.user?.id || undefined,
    // });

    console.error("Fetch error:", error);
    throw new Error(error?.message || "Something went wrong");
  }
};

export const fetchApiSdm = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const baseUrl = `${process.env.NEXT_PUBLIC_API_SDM_URL}/api`;
  const url = `${baseUrl}${endpoint}`;
  // const startTime = Date.now();
  // const method = (options.method || "GET").toUpperCase();
  // const requestId = Math.random().toString(36).substring(7);

  // serverLogger.info(`Server SDM API call: ${method} ${endpoint}`, {
  //   requestId,
  //   method,
  //   endpoint,
  //   url,
  //   hasAuth: !!token,
  //   userId: session?.user?.id,
  // });

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    // const duration = Date.now() - startTime;

    // serverLogger.apiCall({
    //   url: endpoint,
    //   method,
    //   status: response.status,
    //   duration,
    //   requestId,
    //   userId: session?.user?.id || undefined,
    // });

    if (response.status === 401) {
      // serverLogger.warn(`Unauthorized SDM API call: ${method} ${endpoint}`, {
      //   requestId,
      //   status: 401,
      //   duration,
      // });
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();

    // serverLogger.info(`Server SDM API call completed: ${method} ${endpoint}`, {
    //   requestId,
    //   status: response.status,
    //   duration,
    //   success: response.ok,
    // });

    return result;
  } catch (error: any) {
    // const duration = Date.now() - startTime;

    // serverLogger.error(`Server SDM API call failed: ${method} ${endpoint}`, error, {
    //   requestId,
    //   duration,
    //   url,
    //   userId: session?.user?.id || undefined,
    // });

    console.error("Fetch error:", error);
    throw new Error(error?.message || "Something went wrong");
  }
};
