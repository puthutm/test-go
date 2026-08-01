import { Metadata } from "next";
import React from "react";
import CreditLimitTrashClient from "./_components/client";

import { getTrashSksLimit } from "@/services/api/settings/sks-limit/trash";

export const metadata: Metadata = {
  title: "Sampah Batas SKS",
};

export default async function CreditLimitTrashPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {

    const page = Number(searchParams.page) || 1;
    const search = searchParams.search || "";
    const limit = searchParams.limit || 10;
  
    const data = await getTrashSksLimit({
      page: page as number,
      search: search as string,
      limit: limit as number,
    });
  return <CreditLimitTrashClient dataTrashSksLimit={data} />;
}
