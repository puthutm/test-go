import { DefaultSession } from "next-auth";
import { DefaultJWT } from "next-auth/jwt";

declare module "next-auth" {
  /**
   * Returned by `useSession`, `getSession` and received as a prop on the `SessionProvider` React Context
   */
  interface Session {
    user: {
      /** The user's postal address. */
      id?: string | undefined | null;
      email?: string | undefined | null;
      sub?: string | undefined | null;
      app_id?: string | undefined | null;
      role_id?: string | undefined | null;
      role_name?: string | undefined | null;
      exp?: number | undefined | null;
      token?: string | undefined | null;
      refreshToken?: string | undefined | null;
    } & DefaultSession["user"];
  }
}

declare module "next-auth/jwt" {
  interface JWT extends DefaultJWT {
    id?: string | undefined | null;
    email?: string | undefined | null;
    sub?: string | undefined | null;
    app_id?: string | undefined | null;
    role_id?: string | undefined | null;
    role_name?: string | undefined | null;
    expiredIn?: number | undefined | null;
    token?: string | undefined | null;
    refreshToken?: string | undefined | null;
  }
}
