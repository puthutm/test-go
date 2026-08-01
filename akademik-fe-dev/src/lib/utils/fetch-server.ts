"use server";

import { getServerSession } from "next-auth";

import authOptions from "@/config/next-auth";

interface FetchApiOptions extends RequestInit {
  cacheType?: RequestCache;
}

export const fetchApi = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const base = process.env.NEXT_PUBLIC_API_BASE_URL || "http://10.10.20.56:8080";
  const baseUrl = `${base.replace(/\/$/, "")}/api`;
  const url = `${baseUrl}${endpoint.startsWith("/") ? endpoint : `/${endpoint}`}`;

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    if (response.status === 401) {
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();
    return result;
  } catch (error: any) {
    console.error("fetchApi error:", error);
    return { error: true, data: null, message: error?.message || "Service error" };
  }
};

export const fetchApiSso = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const base = process.env.NEXT_PUBLIC_API_SSO_URL || process.env.NEXT_PUBLIC_API_BASE_URL || "http://10.10.20.56:8080";
  const baseUrl = `${base.replace(/\/$/, "")}/api`;
  const url = `${baseUrl}${endpoint.startsWith("/") ? endpoint : `/${endpoint}`}`;

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    if (response.status === 401) {
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();
    return result;
  } catch (error: any) {
    console.error("fetchApiSso error:", error);
    return { error: true, data: [], message: error?.message || "SSO service unavailable" };
  }
};

export const fetchApiDatareferensi = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const base = process.env.NEXT_PUBLIC_API_DATA_REFERENSI_URL || "http://10.10.20.56:3000";
  const baseUrl = `${base.replace(/\/$/, "")}/api`;
  const url = `${baseUrl}${endpoint.startsWith("/") ? endpoint : `/${endpoint}`}`;

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    if (response.status === 401) {
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();
    return result;
  } catch (error: any) {
    console.error("fetchApiDatareferensi error:", error);
    return { error: true, data: null, message: error?.message || "Data Referensi error" };
  }
};

export const fetchApiSdm = async (
  endpoint: string,
  options: FetchApiOptions = {}
): Promise<any> => {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const base = process.env.NEXT_PUBLIC_API_SDM_URL || process.env.NEXT_PUBLIC_API_BASE_URL || "http://10.10.20.56:8080";
  const baseUrl = `${base.replace(/\/$/, "")}/api`;
  const url = `${baseUrl}${endpoint.startsWith("/") ? endpoint : `/${endpoint}`}`;

  try {
    const response = await fetch(url, {
      ...options,
      cache: options.cacheType || "no-cache",
      headers: {
        ...(options.headers || {}),
        Authorization: token ?? "",
      },
    });

    if (response.status === 401) {
      return {
        error: true,
        data: "",
        message: null,
        status: 401,
      };
    }

    const result = await response.json();
    return result;
  } catch (error: any) {
    console.error("fetchApiSdm error:", error);
    return { error: true, data: null, message: error?.message || "SDM service error" };
  }
};
