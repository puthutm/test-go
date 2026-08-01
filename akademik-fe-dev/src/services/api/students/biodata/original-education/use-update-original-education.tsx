"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";

import { FormOriginalEducationSchemaType } from "@/lib/validations/students/biodata/form-original-education-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const useUpdateOriginalEducationStudent = () => {
  const queryClient = useQueryClient();

  const onUpdate = async (
    payload: FormOriginalEducationSchemaType
  ): Promise<ApiResponse<OriginalEducationStudent>> => {
    const reqBody = new FormData();

    if (payload.institution_name)
      reqBody.append("institution_name", payload?.institution_name as string);
    if (payload.school_major)
      reqBody.append("school_major", payload.school_major);
    if (payload.nisn) reqBody.append("nisn", payload.nisn);
    if (payload.national_exam_score)
      reqBody.append("national_exam_score", payload.national_exam_score);
    if (payload.certificate_number)
      reqBody.append("certificate_number", payload.certificate_number);
    if (payload.certificate_filepath)
      reqBody.append("certificate_file", payload.certificate_filepath);
    if (payload.transcripts_filepath)
      reqBody.append("transcripts_file", payload.transcripts_filepath);

    try {
      const response = await fetchApi("/student/biodata/original-educations", {
        method: "PUT",
        body: reqBody,
      });

      queryClient.refetchQueries({
        queryKey: ["original-education-student"],
      });

      return response;
    } catch (error: any) {
      console.log(error?.message, "<<<<<");
      throw new Error(error);
    }
  };

  const mutation = useMutation({
    mutationFn: onUpdate,
  });

  return { ...mutation };
};
