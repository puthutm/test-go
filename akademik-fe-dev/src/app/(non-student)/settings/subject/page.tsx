import { Metadata } from "next";

import { getSubjects } from "@/services/api/settings/subject/get-all-subject";
import { TableSubject } from "./components/table-subject";

export const metadata: Metadata = {
  title: "Mata Kuliah",
  description: "Mata Kuliah",
};
export default async function SubjectPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const data = await getSubjects({
    page: page,
    search: search as string,
  });

  return <TableSubject data={data} />;
}
