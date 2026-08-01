"use server";

import { fetchApiSso } from "@/lib/utils/fetch-server";

export const getProfile = async (): Promise<ApiResponse<Profile>> => {
  try {
    const res = await fetchApiSso("/profile");
    if (!res || res.error) {
      return {
        error: true,
        data: null as any,
        message: res?.message || "Profile unavailable",
      };
    }
    return res;
  } catch (error: any) {
    console.log("Error fetching profile:", error?.message);
    return {
      error: true,
      data: null as any,
      message: error?.message || "Profile unavailable",
    };
  }
};
