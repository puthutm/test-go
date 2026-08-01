import React from "react";
import { Card } from "reactstrap";
import { getServerSession } from "next-auth";
import { Metadata } from "next";

import { getUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/get-unsia-study-program";
import { FormCurriculumStudyProgram } from "./components/form-curriculum-study-program";
import authOptions from "@/config/next-auth";
import { AKADEMIK } from "@/lib/constants/role";
import { getSearchSemesterNumber } from "@/services/api/data-referensi/semester-number/get-search-semester-number";
import { TableCurriculumStudyProgram } from "./components/table-curriculum-study-program";
import { TabStudyProgram } from "./components/tab-study-program-for-academic";

export const metadata: Metadata = {
  title: "Kurikulum Prodi",
  description: "Kurikulum Prodi",
};

export default async function CurriculumStudyProgram({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const session = await getServerSession(authOptions);
  const role = session?.user.role_name;
  const semesters = await getSearchSemesterNumber({ page: 1 });
  let studyProgramParam = searchParams.study_program;
  const curriculumYearParam = searchParams.curriculum_year;

  const unsiaStudyProgram =
    role === AKADEMIK ? await getUnsiaStudyProgram() : null;

  if (!studyProgramParam) {
    studyProgramParam = unsiaStudyProgram?.data?.[0].id as string;
  }

  return (
    <Card>
      <div className="d-flex flex-column gap-3 p-3">
        {role === AKADEMIK ? (
          <TabStudyProgram
            unsiaStudyProgram={
              unsiaStudyProgram as ApiResponse<UnsiaStudyProgram[]>
            }
          />
        ) : null}
        <h2
          className={`${
            role === AKADEMIK ? "mt-2" : "mt-0"
          } mb-0 p-0 fs-4 fw-medium flex-grow-1`}
        >
          Kurikulum Prodi
        </h2>
        <FormCurriculumStudyProgram
          studyProgramId={studyProgramParam as string}
          role={role as string}
        />
        {curriculumYearParam
          ? semesters?.data?.map((data) => (
              <TableCurriculumStudyProgram
                key={data.id}
                semester={data}
                role={role as string}
                studyProgramId={studyProgramParam as string}
              />
            ))
          : null}
      </div>
    </Card>
  );
}
