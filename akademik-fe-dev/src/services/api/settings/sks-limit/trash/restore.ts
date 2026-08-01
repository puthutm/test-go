"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export const useRestoreCreditLimit = () => {
  const axios = useAxios();

  const queryClient = useQueryClient();

  const onRestoreCreditLimit = async (id: string) => {
    try {
      const { data } = await axios.put(
        `/academic/setting/sks-limits/trash/${id}`
      );

      queryClient.invalidateQueries({
        queryKey: ["get-all-credit-limit"],
      });

      queryClient.resetQueries({
        queryKey: ["get-all-trash-credit-limit"],
      });

      queryClient.resetQueries({
        queryKey: ["get-credit-limit", id],
      });

      queryClient.resetQueries({
        queryKey: ["get-search-credit-limit"],
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
    mutationFn: (id: string) => onRestoreCreditLimit(id),
    mutationKey: ["restore-credit-limit"],
  });

  return { ...mutate };
};
