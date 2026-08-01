"use server";

import { fetchApiSso } from "@/lib/utils/fetch-server";

export const getSidebarMenu = async (): Promise<ApiResponse<SidebarMenu[]>> => {
  try {
    const response = await fetchApiSso("/menus/validate/app");

    return response;
  } catch (error: any) {
    console.log(error?.message || "Something went wrong");
    throw new Error(error);
  }
};
