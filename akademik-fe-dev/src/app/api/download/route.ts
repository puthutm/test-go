import { NextRequest, NextResponse } from "next/server";
import { getServerSession } from "next-auth";
import authOptions from "@/config/next-auth";

export async function GET(req: NextRequest) {
  const session = await getServerSession(authOptions);
  const token = session?.user?.token;

  const path = req.nextUrl.searchParams.get("path");
  const baseUrl = `${process.env.NEXT_PUBLIC_API_SSO_URL}/api/objects?path=${path}`;

  if (!path) {
    return NextResponse.json(
      { error: "Missing 'path' query parameter" },
      { status: 400 }
    );
  }

  try {
    const fileRes = await fetch(baseUrl, {
      headers: {
        Authorization: token ?? "",
      },
    });

    if (!fileRes.ok) {
      const errorText = await fileRes.json();
      return NextResponse.json(
        { error: errorText.message },
        { status: fileRes.status }
      );
    }

    const blob = await fileRes.arrayBuffer();

    return new NextResponse(blob, {
      headers: {
        "Content-Type":
          fileRes.headers.get("Content-Type") || "application/octet-stream",
        "Content-Disposition": `attachment; filename="${path
          ?.split("/")
          .pop()}"`,
      },
    });
  } catch (error) {
    return NextResponse.json(
      {
        error: "Internal Server Error",
        message: error instanceof Error ? error.message : "Unknown error",
      },
      { status: 500 }
    );
  }
}
