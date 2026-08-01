import { Metadata } from "next";

import { getCurriculumYears } from "@/services/api/data-referensi/curriculum-year/get-all-curriculum-year";
import { TableCurriculumYear } from "./components/table-curriculum-year";

export const metadata: Metadata = {
  title: "Tahun Kurikulum",
  description: "Tahun Kurikulum",
};

export default async function CurriculumYearPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const data = await getCurriculumYears({
    page: page,
    filter: search as string,
  });

  return <TableCurriculumYear data={data} />;
}
