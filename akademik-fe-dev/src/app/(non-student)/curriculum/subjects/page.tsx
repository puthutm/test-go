import { Metadata } from "next";
import React from "react";
// import SubjectClientPage from "./_components/client";
import PageSubjectLecturer from "./_components/page-subject-lecturer";

// import SubjectClientPage from "./_components/client";
// import ClientTemporary from "./_components/client-temporary";

import { getSubjectLecture } from "@/services/api/curriculum/lecture-subject/get-all-subject-lecture";

export const metadata: Metadata = {
  title: "Mata Kuliah",
};

export default async function SubjectPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.search || "";

  const data = await getSubjectLecture({
    page: page,
    search: search as string,
    curriculum_year_id: searchParams.curriculum_year_id as string,
  });
  return <PageSubjectLecturer dataSubject={data} />;
}
