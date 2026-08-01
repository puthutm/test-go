"use client";

import { Session } from "next-auth";
import { SessionProvider } from "next-auth/react";

const HALF_MINUTES = 30; // IN SECOND

export const AuthProvider = ({
  children,
  session,
}: {
  children: React.ReactNode;
  session: Session;
}) => {
  return (
    <SessionProvider session={session} refetchInterval={HALF_MINUTES}>
      {children}
    </SessionProvider>
  );
};
