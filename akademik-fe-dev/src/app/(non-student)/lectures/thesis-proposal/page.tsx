import React from "react";
import { Metadata } from "next";
import { getFinalProjectProposalForProgramHead } from "@/services/api/program-head/course/final-project-proposal/get-all-final-project-proposal";
import { getServerSession } from "next-auth";
import authOptions from "@/config/next-auth";
import PageClientThesisProposal from "./_components/client";

export const metadata: Metadata = {
  title: "Proposal Tugas Akhir",
};

export default async function PageCollegeClassAcademic({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  let data;
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const academicPeriod = searchParams.academic_period || "";
  const status = searchParams.status || "";
  const studyProgram = searchParams.study_program_id || "";
  const session = await getServerSession(authOptions);
  const role = session?.user.role_name;

  switch (role) {
    case "kaprodi":
      data = await getFinalProjectProposalForProgramHead({
        page: page,
        search: search as string,
        academic_period_id: academicPeriod as string,
        status: status as string,
        study_program_id: studyProgram as string,
      });
      break;
    default:
      data = [] as any;
      break;
  }

  return <PageClientThesisProposal data={data} role={role as string} />;
}
