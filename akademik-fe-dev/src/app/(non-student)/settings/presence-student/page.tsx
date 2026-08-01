import React from "react";
import { Card, Col, Row } from "reactstrap";
import TabsPresences from "./components/tab";
import { TablePresenceStudent } from "./components/table-presence-student";
import { getAllPresenceStudents } from "@/services/api/settings/presence/students/get-all-presence-student";
import { getServerSession } from "next-auth";
import authOptions from "@/config/next-auth";
import { AKADEMIK } from "@/lib/constants/role";
import { notFound } from "next/navigation";

export default async function PresencePage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const session = await getServerSession(authOptions);
  const role = session?.user?.role_name;
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const academicPeriodId = searchParams.academicPeriod || "";
  const studyProgramId = searchParams.studyProgram || "";
  const tabParams = searchParams?.tabs;

  if (role !== AKADEMIK) notFound();

  const data = await getAllPresenceStudents({
    academic_periode_id: academicPeriodId as string,
    study_program_id: studyProgramId as string,
    page: page,
    search: search as string,
  });

  console.log(data?.data);

  return (
    <Row>
      <Col>
        <Card className="pb-4">
          <div className="gap-2 d-flex flex-column w-100 gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Presensi Mahasiswa
            </h1>
            <TabsPresences />
          </div>
          {tabParams === "presence" || !tabParams ? (
            <TablePresenceStudent data={data} />
          ) : (
            <>component</>
          )}
        </Card>
      </Col>
    </Row>
  );
}
