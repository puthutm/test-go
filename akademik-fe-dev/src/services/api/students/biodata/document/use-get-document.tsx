"use client";

import { useQuery } from "@tanstack/react-query";

import { getDocumentStudent } from "./get-document";

export const useDocumentStudent = () => {
  return useQuery({
    queryKey: ["document-student"],
    queryFn: async () => {
      const data = await getDocumentStudent();
      return data;
    },
  });
};
