"use client";

import { useSession } from "next-auth/react";
import { useEffect } from "react";
import { faroEvents, faroSession } from "../faro-utils";

export const useFaroSession = () => {
  const { data: session, status } = useSession();

  useEffect(() => {
    if (status === "authenticated" && session?.user) {
      // Set user information for Faro
      faroSession.setUser(session.user.id || 'unknown', {
        email: session.user.email,
        name: session.user.name,
        role: (session.user as any).role,
      });

      // Track login event
      faroEvents.userLogin(
        session.user.id || 'unknown',
        (session.user as any).role
      );
    } else if (status === "unauthenticated") {
      // Clear user and track logout
      faroSession.clearUser();
      faroEvents.userLogout('unknown');
    }
  }, [session, status]);

  return { session, status };
};