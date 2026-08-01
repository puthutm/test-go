import { Metadata } from "next";
import React from "react";
import PageCollegeClass from "./component/client";

import { getClassScheduleAcademicLecture } from "@/services/api/academic/lecturer/class-schedule/get-class-schedule-academic-lecture";

export const metadata: Metadata = {
  title: "Kelas Kuliah",
};

export default async function SubjectPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.search || "";

  const data = await getClassScheduleAcademicLecture({
    page: page,
    search: search as string,
    curriculum_year_id: searchParams.curriculum_year_id as string,
    study_program_id: searchParams.study_program_id as string,
  });
  return <PageCollegeClass dataClassSchedule={data} />;
}
