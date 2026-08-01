"use server";

import { fetchApiSso } from "@/lib/utils/fetch-server";

export const getSidebarMenu = async (): Promise<ApiResponse<SidebarMenu[]>> => {
  try {
    const response = await fetchApiSso("/menus/validate/app");

    return response;
  } catch (error: any) {
    console.log("Error fetching sidebar menu:", error?.message || "Something went wrong");
    return {
      error: true,
      data: [],
      message: error?.message || "Failed to load sidebar menu",
    } as any;
  }
};
