import { cache } from "react";

import { QueryClient, QueryClientConfig } from "@tanstack/react-query";

const STALE_TIME = 1000 * 60 * 1; // 1 minutes

export const queryClientConfig: QueryClientConfig = {
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      staleTime: STALE_TIME,
      retry: 0,
    },
    mutations: {
      onError: (error: Error) => {
        console.error(error.message || "Something went wrong");
      },
    },
  },
};

export const getQueryClient = cache(() => new QueryClient(queryClientConfig));
