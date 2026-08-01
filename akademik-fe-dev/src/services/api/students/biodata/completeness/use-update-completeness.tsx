"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";

import { FormCompletenessSchemaType } from "@/lib/validations/students/biodata/form-completness-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const useUpdateCompleteness = () => {
  const queryClient = useQueryClient();

  const updateCompletenessStudent = async (
    payload: FormCompletenessSchemaType
  ): Promise<ApiResponse<CompletenessStudent>> => {
    const reqBody = new FormData();

    if (payload.no_passport)
      reqBody.append("no_passport", payload.no_passport as string);
    if (payload.google_scholar)
      reqBody.append("google_scholar", payload.google_scholar as string);
    if (payload.sinta_id)
      reqBody.append("sinta_id", payload.sinta_id as string);
    if (payload.scopus_id)
      reqBody.append("scopus_id", payload.scopus_id as string);
    if (payload.signature_path_file)
      reqBody.append("file_path", payload?.signature_path_file);

    try {
      const response = await fetchApi("/student/biodata/completeness", {
        method: "PUT",
        body: reqBody,
      });

      queryClient.refetchQueries({
        queryKey: ["completeness-student"],
      });

      return response;
    } catch (error: any) {
      console.log(error?.message, "<<<< ERROR");
      throw new Error(error);
    }
  };
  const mutation = useMutation({
    mutationFn: updateCompletenessStudent,
  });

  return { ...mutation };
};
