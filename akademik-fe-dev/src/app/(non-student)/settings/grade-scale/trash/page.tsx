import { Metadata } from "next";
import React from "react";
import GradeScaleTrashPageClient from "./_components/client";

import { getTrashGradeScale } from "@/services/api/settings/grade-scale/trash";

export const metadata: Metadata = {
  title: "Skala Nilai (Sampah)",
};

export default async function GradeScaleTrashPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {

    const page = Number(searchParams.page) || 1;
    const search = searchParams.search || "";
    const limit = searchParams.limit || 10;
  
    const data = await getTrashGradeScale({
      page: page as number,
      search: search as string,
      limit: limit as number,
    });
  

  return <GradeScaleTrashPageClient dataTrashGradeScale={data} />;
}
