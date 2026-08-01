"use server";

import { getServerSession } from "next-auth";

import authOptions from "@/config/next-auth";
import { axiosApi, axiosApiDataReferensi, axiosApiSSO } from "../utils/axios";

export const axiosServerDataReferensi = async () => {
  const session = await getServerSession(authOptions);

  axiosApiDataReferensi.interceptors.request.use((config) => {
    config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
    return config;
  });

  return axiosApiDataReferensi;
};

export const axiosServerSso = async () => {
  const session = await getServerSession(authOptions);

  axiosApiSSO.interceptors.request.use((config) => {
    config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
    return config;
  });

  return axiosApiSSO;
};

export const axiosServer = async () => {
  const session = await getServerSession(authOptions);

  axiosApi.interceptors.request.use((config) => {
    config.headers["Authorization"] = `Bearer ${session?.user?.token}`;
    return config;
  });

  return axiosApi;
};
