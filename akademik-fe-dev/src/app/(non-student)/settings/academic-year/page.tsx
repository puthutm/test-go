import { Metadata } from "next";

import { TableAcademicYear } from "./components/table-academic-year";
import { getAcademicYears } from "@/services/api/data-referensi/academic-year/get-all-academic-year";

export const metadata: Metadata = {
  title: "Tahun Ajaran",
  description: "Tahun Ajaran",
};

export default async function AcademicYearPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.filter || "";
  const data = await getAcademicYears({
    page,
    filter: search as string,
  });
  return <TableAcademicYear data={data} />;
}
