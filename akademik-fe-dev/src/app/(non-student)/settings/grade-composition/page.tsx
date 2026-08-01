import { Metadata } from "next";
import React from "react";
import GradeCompositionClientPage from "./_components/client";

import { getAllGradeCompositions } from "@/services/api/settings/grade-composition/get-all-grade-composition";

export const metadata: Metadata = {
  title: "Komposisi Nilai",
  description:
    "Komposisi Nilai adalah halaman yang digunakan untuk mengatur komposisi nilai yang digunakan dalam penilaian.",
};

export default async function GradeCompositionPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
    const page = Number(searchParams.page) || 1;
    const search = searchParams.search || "";
    const limit = searchParams.limit || 10;
  
    const data = await getAllGradeCompositions({
      page: page as number,
      search: search as string,
      limit: limit as number,
      value_element_id: searchParams.value_element_id as string,
    });
  return <GradeCompositionClientPage dataGradeComposition={data} />;
}
