"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormOpenCloseGradeSchemaType } from "@/lib/validations/settings/academic-period/form-open-close-grade";
import { revalidatePath } from "next/cache";

interface Args {
  academicPeriodId: string;
  classId: string;
  payload: FormOpenCloseGradeSchemaType;
}

export const updateOpenCloseGradeByClassId = async ({
  academicPeriodId,
  classId,
  payload,
}: Args): Promise<ApiResponse<OpenCloseGrade> | undefined> => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/${academicPeriodId}/classes/${classId}/class-score/open-close-values`,
      {
        method: "PUT",
        body: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath(
      "/settings/academic-period/[academicPeriodId]/classes/[classId]/detail",
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
