"use client";

// import { signOut } from "next-auth/react";
import { useFaroSession } from "../hooks/use-faro-session";

export const useFetchClient = () => {
  const { session } = useFaroSession();

  const fetchClient = async (
    endpoint: string
  ): Promise<Response | undefined> => {
    try {
      const api = await fetch(
        `${process.env.NEXT_PUBLIC_API_BASE_URL}/api${endpoint}`,
        {
          cache: "no-cache",
          headers: {
            Authorization: `Bearer ${session?.user.token as string}`,
          },
        }
      );

      // if (api.status === 401) {
      //   await signOut();
      // }

      if (!api.ok) {
        throw new Error("Failed to fetch data");
      }

      return api;
    } catch (error) {
      if (error instanceof Error) throw new Error(error.message);
    }
  };

  const fetchSsoClient = async (
    endpoint: string
  ): Promise<Response | undefined> => {
    try {
      const api = await fetch(
        `${process.env.NEXT_PUBLIC_API_SSO_URL}/api${endpoint}`,
        {
          cache: "no-cache",
          headers: {
            Authorization: `Bearer ${session?.user.token as string}`,
          },
        }
      );

      // if (api.status === 401) {
      //   await signOut();
      // }

      if (!api.ok) {
        throw new Error("Failed to fetch data");
      }

      return api;
    } catch (error) {
      if (error instanceof Error) throw new Error(error.message);
    }
  };

  const fetchDataReferensiClient = async (
    endpoint: string
  ): Promise<Response | undefined> => {
    try {
      const api = await fetch(
        `${process.env.NEXT_PUBLIC_API_DATA_REFERENSI_URL}/api${endpoint}`,
        {
          cache: "no-cache",
          headers: {
            Authorization: `Bearer ${session?.user.token as string}`,
          },
        }
      );

      // if (api.status === 401) {
      //   await signOut();
      // }

      if (!api.ok) {
        throw new Error("Failed to fetch data");
      }

      return api;
    } catch (error) {
      if (error instanceof Error) throw new Error(error.message);
    }
  };

  const fetchSdmClient = async (
    endpoint: string
  ): Promise<Response | undefined> => {
    try {
      const api = await fetch(
        `${process.env.NEXT_PUBLIC_API_SDM_URL}/api${endpoint}`,
        {
          cache: "no-cache",
          headers: {
            Authorization: `Bearer ${session?.user.token as string}`,
          },
        }
      );

      // if (api.status === 401) {
      //   await signOut();
      // }

      if (!api.ok) {
        throw new Error("Failed to fetch data");
      }

      return api;
    } catch (error) {
      if (error instanceof Error) throw new Error(error.message);
    }
  };

  return {
    fetchClient,
    fetchSsoClient,
    fetchDataReferensiClient,
    fetchSdmClient,
  };
};
