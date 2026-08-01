import { getToken } from "next-auth/jwt";
import { NextRequest, NextResponse } from "next/server";
import { MAHASISWA } from "./lib/constants/role";

const middleware = async (request: NextRequest) => {
  const token = await getToken({
    req: request,
    secret: process.env.NEXTAUTH_SECRET,
  });

  const isAuthenticated = !!token;
  const role = token?.role_name;
  const path = request.nextUrl.pathname;

  if (isAuthenticated && path === "/") {
    if (role === MAHASISWA) {
      return NextResponse.redirect(new URL("/student", request.url));
    }
    return NextResponse.redirect(new URL("/dashboard", request.url));
  }

  if (isAuthenticated && path === "/dashboard" && role === MAHASISWA) {
    return NextResponse.redirect(new URL("/student", request.url));
  }

  if (!isAuthenticated && path !== "/") {
    return NextResponse.redirect(new URL("/", request.url));
  }

  return NextResponse.next();
};

export default middleware;

export const config = {
  matcher: [
    "/((?!api/auth|api|metrics|_next/static|_next/image|manifest.webmanifest|favicon.ico|sitemap.xml|robots.txt|.*\\.png|.*\\.jpg|.*\\.jpeg$).*)",
  ],
};
