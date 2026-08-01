"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";

import { FormBankAccountSchemaType } from "@/lib/validations/students/biodata/form-bank-account-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const useUpdateBankAccountStudent = () => {
  const queryClient = useQueryClient();

  const onUpdate = async (
    payload: FormBankAccountSchemaType
  ): Promise<ApiResponse<BankAccountStudent>> => {
    const reqBody = new FormData();

    if (payload.bank_id)
      reqBody.append("bank_id", payload?.bank_id?.value as string);
    if (payload.account_number)
      reqBody.append("account_number", payload.account_number);
    if (payload.account_name)
      reqBody.append("account_name", payload.account_name);
    if (payload.account_file_path)
      reqBody.append("account_file_path", payload.account_file_path);

    try {
      const response = await fetchApi("/student/biodata/bank-accounts", {
        method: "PUT",
        body: reqBody,
      });

      queryClient.refetchQueries({
        queryKey: ["bank-account-student"],
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
