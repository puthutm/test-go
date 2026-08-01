import { Metadata } from "next";
import React from "react";
import GradeScaleClientPage from "./_components/client";

import { getGradeScale } from "@/services/api/settings/grade-scale/get-all-grade-scale";
export const metadata: Metadata = {
  title: "Skala Nilai",
  description:
    "Skala Nilai adalah halaman yang digunakan untuk mengatur skala nilai yang digunakan dalam penilaian.",
};

export default async function GradeScalePage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.search || "";
  const limit = searchParams.limit || 10;

  const data = await getGradeScale({
    page: page as number,
    search: search as string,
    limit: limit as number,
    study_program_id: searchParams.study_program_id as string,
  });

  return <GradeScaleClientPage dataGradeScale={data} />;
}
