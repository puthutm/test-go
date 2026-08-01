"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useRestoreSubject = () => {
  const axios = useAxios();
  const queryClient = useQueryClient();

  const onRestoreSubject = async (id: string) => {
    try {
      const { data } = await axios.put(
        `/academic/setting/subjects/trash/${id}`
      );

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
        error instanceof Error
          ? error.message
          : "An error occured while restoring data"
      );
    }
  };

  const mutate = useMutation({
    mutationFn: (id: string) => onRestoreSubject(id),
    mutationKey: ["restore-subject"],
  });

  return { ...mutate };
};
