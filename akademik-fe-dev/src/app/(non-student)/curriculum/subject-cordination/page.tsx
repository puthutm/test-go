import { Metadata } from "next";
import React from "react";
import { getSubjectLectureCordination } from "@/services/api/curriculum/lecture-subject-cordinator/get-all-subject-cordinator";
import PageSubjectCordination from "./components/client";

export const metadata: Metadata = {
  title: "Mata Kuliah Kordinator",
};

export default async function SubjectPage({
  searchParams
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {

  const page = Number(searchParams.page) || 1;
  const search = searchParams.search || "";

  const data = await getSubjectLectureCordination({
    page: page,
    search: search as string,
    curriculum_year_id:searchParams.curriculum_year_id as string
  });
  return <PageSubjectCordination dataSubjectCordination={data}/>
}
