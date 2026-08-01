import { Metadata } from "next";
import React from "react";
import GradeCompositionTrashPageClient from "./_components/client";

import { getTrashGradeComposition } from "@/services/api/settings/grade-composition/trash";

export const metadata: Metadata = {
  title: "Komposisi Nilai (Sampah)",
};

export default async function GradeCompositionTrashPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {

  
      const page = Number(searchParams.page) || 1;
      const search = searchParams.search || "";
      const limit = searchParams.limit || 10;
    
      const data = await getTrashGradeComposition({
        page: page as number,
        search: search as string,
        limit: limit as number,
      });

  return <GradeCompositionTrashPageClient  dataTrashGradeComposition={data} />;
}
