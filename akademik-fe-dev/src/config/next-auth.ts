import { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

import { decodeJWT } from "@/lib/utils/jwt-decode";
import { getRefreshAccessToken } from "@/services/api/auth/refresh-token";

const FOUR_MINUTES = 240000; //IN SECOND

const authOptions: NextAuthOptions = {
  pages: {
    signIn: "/",
  },
  session: {
    strategy: "jwt",
    maxAge: 3600, //1 hours
  },
  secret: process.env.NEXTAUTH_SECRET,
  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        username: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
        role: { label: "Role", type: "text" },
        app: { label: "App", type: "text" },
        token: { label: "Token", type: "password" },
      },
      async authorize(credentials: any) {
        try {
          const { username, role, token, app } = credentials as {
            username?: string;
            password?: string;
            role?: string;
            token?: string;
            app?: string;
          };

          // Local direct login mode
          if (username || (role && !token)) {
            const roleName = role || "mahasiswa";
            const user = {
              id: username || "user-local-id",
              email: `${username || "user"}@unsia.ac.id`,
              sub: username || "user-sub",
              app_id: app || "app-akademik",
              role_id: "role-local-id",
              role_name: roleName,
              exp: Math.floor(Date.now() / 1000) + 3600 * 24,
              token: "local-dummy-token",
              refreshToken: "local-dummy-refresh-token",
            };
            return Promise.resolve(user);
          }

          // SSO credentials fetch
          const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/auth/role-spesifik`,
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer ${token}`,
              },
              body: JSON.stringify({
                role_id: role,
                app_id: app,
              }),
            }
          );

          const { data, message } = await response.json();

          if (!data) {
            throw new Error(message || "Something went wrong");
          }

          const dataToken = decodeJWT(data?.token);

          const user = {
            id: data?.id || dataToken?.id,
            email: dataToken?.email,
            sub: dataToken?.sub,
            app_id: dataToken?.app_id,
            role_id: dataToken?.role_id,
            role_name: dataToken?.role_name,
            exp: dataToken?.exp,
            token: data?.token,
            refreshToken: data?.refresh_token,
          };
          return Promise.resolve(user);
        } catch (err: any) {
          throw new Error(err);
        }
      },
    }),
    {
      id: "sso",
      name: "sso",
      type: "oauth",
      version: "2.0",
      authorization: {
        url: `${process.env.NEXT_PUBLIC_API_SSO_URL}/oauth/authorization`,
        params: { scope: "" },
      },
      checks: ["state"],
      token: {
        url: `${process.env.NEXT_PUBLIC_API_SSO_URL}/oauth/access_token`,
        request: async (context: any) => {
          const requestBody = {
            grant_type: "authorization_code",
            client_id: context.provider.clientId!,
            client_secret: context.provider.clientSecret!,
            code: context.params?.code || context.params.code,
            redirect_uri: `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/auth/callback/sso`,
          };

          const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_SSO_URL}/oauth/access_token`,
            {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify(requestBody),
            }
          );

          const data = await response.json();

          const tokens = {
            access_token: data.data?.AccessToken,
            refresh_token: data.data?.RefreshToken,
            token_type: data.data?.TokenType || "Bearer",
            expires_in: 3600,
          };

          return { tokens };
        },
      },
      userinfo: {
        url: `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/users/profile`,
        request: async (context: any) => {
          const response = await fetch(
            `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/profile`,
            {
              headers: {
                Authorization: `Bearer ${context.tokens.access_token}`,
                "Content-Type": "application/json",
              },
            }
          );

          const data = await response.json();

          return data;
        },
      },
      clientId: process.env.OAUTH_SSO_CLIENT_ID,
      clientSecret: process.env.OAUTH_SSO_CLIENT_SECRET,

      profile(profile: any, tokens: any) {
        return {
          id: profile.data.nik,
          name: profile.data.name,
          username: profile.data.username,
          email: profile.data.email,
          roles: profile.data.roles,
          accessToken: tokens?.access_token,
          refreshToken: tokens?.refresh_token,
        };
      },
      httpOptions: {
        timeout: 10000,
      },
    },
  ],
  callbacks: {
    async jwt({ token, user, account }: any) {
      if (account?.provider === "sso") {
        const dataToken = decodeJWT(account.access_token);

        return {
          id: dataToken?.id,
          email: dataToken?.email,
          role_name: dataToken?.role_name,
          expiredIn: (dataToken?.exp as number) * 1000,
          token: account.access_token,
          refreshToken: account.refresh_token,
        };
      }

      // check is user exist
      if (user) {
        token.id = user?.id;
        token.email = user?.email;
        token.app_id = user?.app_id;
        token.role_id = user?.role_id;
        token.role_name = user?.role_name;
        token.token = user?.token;
        token.expiredIn = user?.exp * 1000;
        token.refreshToken = user?.refreshToken;
      }

      const expDate = token.expiredIn - FOUR_MINUTES;
      // check if token not expired
      if (Date.now() < expDate) {
        return token;
      }

      // if expired refresh token
      return getRefreshAccessToken(token);
    },
    async session({ session, token }: any) {
      if (token) {
        session.user = {};
        session.user.id = token.id;
        session.user.email = token.email;
        session.user.sub = token.sub;
        session.user.app_id = token.app_id;
        session.user.role_id = token.role_id;
        session.user.role_name = token.role_name;
        session.user.token = token.token;
        session.user.exp = token.expiredIn;
        session.user.refreshToken = token.refreshToken;
      }
      return session;
    },
    async redirect({ url, baseUrl }: any) {
      const redirectUrl = url.startsWith("/")
        ? new URL(url, baseUrl).toString()
        : url;
      return redirectUrl;
    },
  },
  debug: process.env.NODE_ENV === "development" ? true : false,
};

export default authOptions;
