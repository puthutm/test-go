"use client";

import { useQuery } from "@tanstack/react-query";
import { getPresenceComponentBySessionId } from "./get-presence-component-by-session-id";

export const useGetPresenceComponentBySessionId = ({
  sessionId,
}: {
  sessionId: string;
}) => {
  return useQuery({
    queryKey: ["presence-component", sessionId],
    queryFn: async () => await getPresenceComponentBySessionId({ sessionId }),
    enabled: !!sessionId,
  });
};
