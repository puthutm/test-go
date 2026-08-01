"use server";

import { fetchApiSso } from "@/lib/utils/fetch-server";

export const getProfile = async (): Promise<ApiResponse<Profile>> => {
  try {
    return await fetchApiSso("/profile");
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
