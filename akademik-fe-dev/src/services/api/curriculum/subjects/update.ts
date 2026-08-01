"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { SubjectFormType } from "@/lib/validations/curriculum/subject";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useUpdateSubject = (id: string) => {
  const axios = useAxios();
  const queryClient = useQueryClient();

  const onUpdateSubject = async (form: SubjectFormType) => {
    try {
      const reqBody = { ...form };

      const { data } = await axios.put(
        `/academic/setting/subjects/${id}`,
        reqBody
      );

      queryClient.invalidateQueries({
        queryKey: ["get-all-subjects"],
      });

      queryClient.resetQueries({
        queryKey: ["get-all-subjects-trash"],
      });

      queryClient.resetQueries({
        queryKey: ["get-subject", id],
      });

      queryClient.resetQueries({
        queryKey: ["get-search-subject"],
      });

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error
          ? error.message
          : "An error occurred while fetching data"
      );
    }
  };

  const mutate = useMutation({
    mutationFn: onUpdateSubject,
    mutationKey: ["update-subject", id],
  });

  return { ...mutate };
};
