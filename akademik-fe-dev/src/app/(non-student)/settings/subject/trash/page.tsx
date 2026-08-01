import React from "react";

import { TableSubjectTrash } from "./components/table-subject-trash";
import { getSubjectsTrash } from "@/services/api/settings/subject/trash/get-all-trash-subject";

export default async function SubjectPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const data = await getSubjectsTrash({
    page: page,
    search: search as string,
  });

  return <TableSubjectTrash data={data} />;
}
