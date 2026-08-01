"use client";

import { useQuery } from "@tanstack/react-query";

import { getBankAccountStudent } from "./get-bank-account";

export const useBankAccountStudent = () => {
  return useQuery({
    queryKey: ["bank-account-student"],
    queryFn: async () => {
      const data = await getBankAccountStudent();
      return data;
    },
  });
};
