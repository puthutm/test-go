import { redirect } from "next/navigation";
import { decodeJWT } from "@/lib/utils/jwt-decode";

export const getRefreshAccessToken = async (token: any) => {
  if (
    !token?.refreshToken ||
    token?.refreshToken === "local-dummy-refresh-token" ||
    token?.token === "local-dummy-token"
  ) {
    return {
      ...token,
      expiredIn: Date.now() + 3600 * 24 * 1000,
    };
  }

  const req = {
    refresh_token: token?.refreshToken,
    role_spesifik: token?.role_name,
    app_id: token?.app_id,
  };

  try {
    const api = await fetch(
      `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/auth/token`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(req),
      }
    );

    const response = await api.json();

    if (response.status === 401 || response.data === null) {
      redirect(process.env.NEXT_PUBLIC_UI_SSO_URL as string);
    }

    const dataToken = decodeJWT(response?.data?.token);
    const user = {
      id: response?.data?.id || dataToken?.id,
      email: dataToken?.email,
      sub: dataToken?.sub,
      app_id: dataToken?.app_id,
      role_id: dataToken?.role_id,
      role_name: dataToken?.role_name,
      exp: dataToken?.exp,
      token: response?.data?.token,
      refreshToken: response?.data?.refresh_token,
    };
    return user;
  } catch (error: any) {
    console.error("Refresh token error:", error);
    return {
      ...token,
      expiredIn: Date.now() + 3600 * 24 * 1000,
    };
  }
};
