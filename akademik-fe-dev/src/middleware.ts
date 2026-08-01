import { getToken } from "next-auth/jwt";
import { NextRequest } from "next/server";
import { NextResponse } from "next/server";
import { MAHASISWA } from "./lib/constants/role";

const middleware = async (request: NextRequest) => {
  const sessionCookie =
    process.env.NODE_ENV === "development"
      ? request.cookies.get("next-auth.session-token")
      : request.cookies.get("__Secure-next-auth.session-token");

  const token = await getToken({
    req: request,
    secret: process.env.NEXTAUTH_SECRET,
  });

  const role = token?.role_name;

  const path = request.nextUrl.pathname;

  if (sessionCookie && path === "/" && role === MAHASISWA) {
    return NextResponse.redirect(new URL("/student", request.url));
  }

  if (sessionCookie && path === "/dashboard" && role === MAHASISWA) {
    return NextResponse.redirect(new URL("/student", request.url));
  }

  if (sessionCookie && path === "/") {
    return NextResponse.redirect(new URL("/dashboard", request.url));
  }

  if (!sessionCookie && path === "/") {
    return NextResponse.next();
  }

  if (!sessionCookie && path !== "/") {
    return NextResponse.redirect(new URL("/", request.url));
  }

  //

  return NextResponse.next();
};

export default middleware;

// ignore middleware
export const config = {
  matcher: [
    "/((?!api/auth|api|metrics|_next/static|_next/image|manifest.webmanifest|favicon.ico|sitemap.xml|robots.txt|.*\\.png|.*\\.jpg|.*\\.jpeg$).*)",
  ],
};
