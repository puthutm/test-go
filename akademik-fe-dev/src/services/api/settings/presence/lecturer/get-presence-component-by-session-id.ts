"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getPresenceComponentBySessionId = async ({
  sessionId,
}: {
  sessionId: string;
}): Promise<ApiResponse<PresenceComponent> | undefined> => {
  try {
    const response = await fetchApi(
      `/lecturer/academic/presence/students/components/sessions/${sessionId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
