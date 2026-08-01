import { jwtDecode, JwtPayload } from "jwt-decode";

interface Users extends JwtPayload {
  id: string | undefined;
  email: string | undefined;
  app_id?: string | undefined | null;
  role_id?: string[] | undefined | null;
  role_name?: string[] | undefined | null;
  sub?: string | undefined;
  exp: number | undefined;
  iat?: number | undefined;
  jti?: string | undefined;
}

export function decodeJWT(token: string): Users | undefined {
  try {
    const decoded = jwtDecode(token) as Users;
    return decoded;
  } catch (err: any) {
    throw new Error(err);
  }
}
