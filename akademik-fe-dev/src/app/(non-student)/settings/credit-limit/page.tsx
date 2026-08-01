import { Metadata } from "next";
import React from "react";
import CreditLimitClientPage from "./_components/client";

import { getAllSksLimits } from "@/services/api/settings/sks-limit/get-all-sks-limits";

export const metadata: Metadata = {
  title: "Batas Sks",
  description:
    "Batas Kredit adalah halaman yang digunakan untuk mengatur batas kredit yang digunakan dalam penilaian.",
};

export default async function CreditLimitPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
    const page = Number(searchParams.page) || 1;
    const search = searchParams.search || "";
    const limit = searchParams.limit || 10;
  
    const data = await getAllSksLimits({
      page: page as number,
      search: search as string,
      limit: limit as number,
    });
  return <CreditLimitClientPage dataSksLimit={data} />;
}
