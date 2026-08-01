import React from "react";
import AcademicPeriodClient from "./components/client";
import { Metadata } from "next";
import { getAcademicPeriods } from "@/services/api/data-referensi/academic-period/get-all-academic-period";

export const metadata: Metadata = {
  title: "Periode Akademik",
  description: "Periode Akademik",
};

export default async function AcademicPeriodPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.filter || "";

  const data = await getAcademicPeriods({
    page: page,
    filter: search as string,
  });

  return <AcademicPeriodClient data={data} />;
}
