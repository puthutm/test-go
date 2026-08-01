"use client";

import { useSession } from "next-auth/react";
import { useEffect } from "react";

import { axiosApi, axiosApiDataReferensi, axiosApiSSO } from "../utils/axios";

export const useAxios = () => {
  const { data: session } = useSession();
  useEffect(() => {
    const axiosIntercept = axiosApi.interceptors.request.use((config) => {
      config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
      return config;
    });
    return () => {
      axiosApi.interceptors.request.eject(axiosIntercept);
    };
  }, [session]);

  return axiosApi;
};

export const useAxiosSSO = () => {
  const { data: session } = useSession();
  useEffect(() => {
    const axiosIntercept = axiosApiSSO.interceptors.request.use((config) => {
      config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
      return config;
    });
    return () => {
      axiosApiSSO.interceptors.request.eject(axiosIntercept);
    };
  }, [session]);

  return axiosApiSSO;
};

export const useAxiosDataReferensi = () => {
  const { data: session } = useSession();
  useEffect(() => {
    const axiosIntercept = axiosApiDataReferensi.interceptors.request.use(
      (config) => {
        config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
        return config;
      }
    );
    return () => {
      axiosApiDataReferensi.interceptors.request.eject(axiosIntercept);
    };
  }, [session]);

  return axiosApiDataReferensi;
};
