"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";

import { FormDocumentSchemaType } from "@/lib/validations/students/biodata/form-document-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const useUpdateDocumentStudent = () => {
  const queryClient = useQueryClient();

  const onUpdate = async (
    payload: FormDocumentSchemaType
  ): Promise<ApiResponse<DocumentStudent>> => {
    const reqBody = new FormData();

    if (payload.npwp) reqBody.append("npwp", payload.npwp);
    if (payload.npwp_filepath)
      reqBody.append("npwp_filepath", payload.npwp_filepath);
    if (payload.bpjs_healthcare)
      reqBody.append("bpjs_healthcare", payload.bpjs_healthcare);
    if (payload.bpjs_healthcare_filepath)
      reqBody.append(
        "bpjs_healthcare_filepath",
        payload.bpjs_healthcare_filepath
      );
    if (payload.bpjs_employment)
      reqBody.append("bpjs_employment", payload.bpjs_employment);
    if (payload.bpjs_employment_filepath)
      reqBody.append(
        "bpjs_employment_filepath",
        payload.bpjs_employment_filepath
      );

    try {
      const response = await fetchApi("/student/biodata/documents", {
        method: "PUT",
        body: reqBody,
      });

      queryClient.refetchQueries({
        queryKey: ["document-student"],
      });

      return response;
    } catch (error: any) {
      console.log(error?.message, "<<<< ERROR");
      throw new Error(error);
    }
  };

  const mutation = useMutation({
    mutationFn: onUpdate,
  });

  return { ...mutation };
};
