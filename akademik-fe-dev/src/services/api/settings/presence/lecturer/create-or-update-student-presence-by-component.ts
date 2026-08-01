"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

interface Payload {
  presenceStatus: string;
  studentId: string;
  presenceType: StudentPresenceComponent;
}

export const createOrUpdateStudentPresenceByComponent = async ({
  payload,
  sessionId,
}: {
  sessionId: string;
  payload: Payload;
}) => {
  try {
    const reqBody = {
      student_id: payload.studentId,
      presence_status: payload.presenceStatus === "true" ? true : false,
      presence_type: payload.presenceType,
    };

    const response = await fetchApi(
      `/lecturer/academic/presence/students/sessions/${sessionId}`,
      {
        method: "POST",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath(
      "/settings/presence/[subjectId]/clasess/[classId]/sessions/[sessionId]/students",
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
