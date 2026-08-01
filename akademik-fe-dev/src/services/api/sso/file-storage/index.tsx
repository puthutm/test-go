"use client";

import { useQuery } from "@tanstack/react-query";

import { useFetchClient } from "@/lib/utils/fetch-client";

export const useGetFileStorage = (src: string) => {
  const { fetchSsoClient } = useFetchClient();

  const fetchFileStorage = async () => {
    try {
      const api = await fetchSsoClient(`/objects?path=${src}`);

      if (!api?.ok) {
        throw new Error("Failed to fetch data");
      }

      const data = await api.blob();

      return data;
    } catch (error) {
      if (error instanceof Error) throw new Error(error.message);
    }
  };

  const query = useQuery({
    queryFn: fetchFileStorage,
    queryKey: ["file-storage", src],
    enabled: !!src,
  });

  return { ...query };
};
