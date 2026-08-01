"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useDeleteSubject = () => {
  const axios = useAxios();
  const queryClient = useQueryClient();

  const deleteSubject = async (id: string) => {
    try {
      const { data } = await axios.delete(`/academic/setting/subjects/${id}`);

      queryClient.invalidateQueries({
        queryKey: ["get-all-subject"],
      });

      queryClient.resetQueries({
        queryKey: ["get-all-subject-trash"],
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
        error instanceof Error ? error.message : "Something went wrong"
      );
    }
  };

  const mutate = useMutation({
    mutationKey: ["delete-subject"],
    mutationFn: (id: string) => deleteSubject(id),
  });

  return { ...mutate };
};
