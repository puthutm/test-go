"use server";

import { revalidatePath } from "next/cache";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";
import { AcademicPeriodFormType } from "@/lib/validations/academic/settings/academic-period";
import { formatDateNumeric } from "@/lib/utils/format-date";

export const createAcademicPeriod = async (payload: AcademicPeriodFormType) => {
  const reqBody = {
    code: payload.code,
    fullname: payload.fullname,
    shortname: payload.shortname,
    start_date_of_college: formatDateNumeric(
      payload.start_date_of_college[0].toString()
    ),
    end_date_of_college: formatDateNumeric(
      payload.end_date_of_college[0].toString()
    ),
    //   payload.start_date_of_uas[0].toString()
    // ),
    // end_date_of_uas: formatDateNumeric(payload.end_date_of_uas[0].toString()),
    // start_date_of_uts: formatDateNumeric(
    //   payload.start_date_of_uts[0].toString()
    // ),
    // end_date_of_uts: formatDateNumeric(payload.end_date_of_uts[0].toString()),
    // is_active: payload.is_active,
    number_of_lecture_meeting: payload.number_of_lecture_meeting.value,
    academic_year_id: payload.academic_year_id.value,
    semester_id: payload.semester_id.value,
  };

  try {
    const response = await fetchApiDatareferensi("/pmb/academic-periods", {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/academic-period", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
