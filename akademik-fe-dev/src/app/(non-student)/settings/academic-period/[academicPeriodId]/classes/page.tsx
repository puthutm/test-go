import { Metadata } from "next";
import { Card } from "reactstrap";
import Link from "next/link";

import { ClassTable } from "./components/class-table";
import AcademicPeriodInfo from "@/components/ui/academic-period-info";
import { getUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/get-unsia-study-program";
import { getAllClassByAcademicPeriodIdForAcademic } from "@/services/api/settings/academic-period/class/get-all-class";

export const metadata: Metadata = {
  title: "Daftar Kelas",
  description: "Daftar Kelas",
};

export default async function AcademicPeriodDetailPage({
  params,
  searchParams,
}: {
  params: Promise<{ academicPeriodId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const academicPeriodId = (await params).academicPeriodId;
  const unsiaStudyProgram = await getUnsiaStudyProgram();

  if (!searchParams.tabs) {
    searchParams.tabs = unsiaStudyProgram?.data?.[0].id as string;
  }

  const data = await getAllClassByAcademicPeriodIdForAcademic({
    academic_periode_id: academicPeriodId as string,
    study_program_id: searchParams.tabs as string,
    page: page,
    search: search as string,
  });

  return (
    <Card>
      <div className="d-flex flex-column gap-3 p-3">
        <AcademicPeriodInfo params={params} />
        <div className="row gap-3 flex-wrap border-bottom pb-3 px-3">
          {unsiaStudyProgram?.data?.map((tab) => (
            <Link
              href={`/settings/academic-period/${academicPeriodId}/classes?tabs=${tab.id}`}
              key={tab.id}
              className={`col rounded-top text-center py-2 px-4 fw-semibold ${
                searchParams.tabs === tab.id ? "bg-primary text-white" : ""
              }`}
              style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
            >
              {tab.name}
            </Link>
          ))}
        </div>
        <ClassTable searchParams={searchParams} data={data} />
      </div>
    </Card>
  );
}
