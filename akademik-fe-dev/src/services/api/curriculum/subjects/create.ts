"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { SubjectFormType } from "@/lib/validations/curriculum/subject";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useCreateSubject = () => {
  const axios = useAxios();
  const queryClient = useQueryClient();

  const onCreateSubject = async (form: SubjectFormType) => {
    try {
      const reqBody = {
        ...form,
      };

      const { data } = await axios.post("/academic/setting/subjects", reqBody);

      queryClient.invalidateQueries({
        queryKey: ["get-all-subjects"],
      });

      queryClient.resetQueries({
        queryKey: ["get-all-subjects-trash"],
      });

      queryClient.resetQueries({
        queryKey: ["get-search-subject"],
      });

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error
          ? error.message
          : "An error occured while fetching subject data"
      );
    }
  };

  const mutate = useMutation({
    mutationFn: onCreateSubject,
    mutationKey: ["create-subject"],
  });

  return { ...mutate };
};
