import axios from "axios";

export const axiosApi = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_API_BASE_URL}/api`,
});

export const axiosApiSSO = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_API_SSO_URL}/api`,
});

export const axiosApiDataReferensi = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_API_DATA_REFERENSI_URL}/api`,
});
