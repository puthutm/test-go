"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetSubject = (id?: string | null) => {
  const axios = useAxios();

  const fetchDetailSubject = async (): Promise<
    ApiResponse<Subject | null> | undefined
  > => {
    try {
      if (!id) {
        return {
          data: null,
          error: true,
          message: "ID is not provided",
          status: 404,
        };
      }

      const { data } = await axios.get(`/academic/setting/subjects/${id}`);

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error ? error.message : "Something went wrong"
      );
    }
  };

  const query = useQuery({
    queryKey: ["get-subject", id],
    queryFn: fetchDetailSubject,
  });

  return { ...query };
};
