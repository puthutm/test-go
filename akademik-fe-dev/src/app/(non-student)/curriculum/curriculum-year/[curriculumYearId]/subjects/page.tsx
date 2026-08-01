import React from "react";
import { Card } from "reactstrap";

import { getUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/get-unsia-study-program";
import Link from "next/link";
import CurriculumYearInfo from "@/components/ui/curriculum-year-info";
import { TableSubject } from "./components/table-subject";
import { getSubjectsByCurriculumYearId } from "@/services/api/settings/subject/get-all-subject-by-curriculum-year-id";

export default async function SubjectPageByCurriculumIdPage({
  params,
  searchParams,
}: {
  params: Promise<{ curriculumYearId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const studyProgram = searchParams.tabs || 0;

  const curriculumYearId = (await params).curriculumYearId;

  const unsiaStudyProgram = await getUnsiaStudyProgram();

  const data = await getSubjectsByCurriculumYearId({
    curriculumYearId,
    queryParam: {
      page,
      search: search as string,
      study_program_id: studyProgram as string,
    },
  });

  return (
    <Card>
      <div className="d-flex flex-column gap-3 p-3">
        <CurriculumYearInfo params={params} />
        <div className="row gap-3 flex-wrap border-bottom pb-2 mx-1">
          <Link
            href={`/settings/curriculum-year/${curriculumYearId}/subject`}
            className={`col rounded-top text-center py-2 px-4 fw-semibold ${
              !searchParams.tabs ? "bg-primary text-white" : ""
            }`}
            style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
            prefetch
          >
            Universitas
          </Link>
          {unsiaStudyProgram?.data?.map((tab) => (
            <Link
              href={`/settings/curriculum-year/${curriculumYearId}/subject?tabs=${tab.id}`}
              key={tab.id}
              className={`col rounded-top text-center py-2 px-4 fw-semibold ${
                searchParams.tabs === tab.id ? "bg-primary text-white" : ""
              }`}
              style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
              prefetch
            >
              {tab.name}
            </Link>
          ))}
        </div>
        <TableSubject data={data} />
      </div>
    </Card>
  );
}
