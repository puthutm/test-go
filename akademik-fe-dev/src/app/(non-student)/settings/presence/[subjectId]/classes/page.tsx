import React from "react";
import { Card, Col, Row } from "reactstrap";
import { getServerSession } from "next-auth";
import { notFound, redirect } from "next/navigation";

import authOptions from "@/config/next-auth";
import { DOSEN } from "@/lib/constants/role";
import { TablePresenceClass } from "./components/table-presence-class";
import { getAllPresenceClassBySubjectId } from "@/services/api/settings/presence/lecturer/get-all-presence-class";
import { getAllPresenceComponent } from "@/services/api/settings/presence/lecturer/get-presence-component";

export default async function PresencePage({
  searchParams,
  params,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  params: { subjectId: string };
}) {
  const session = await getServerSession(authOptions);
  const role = session?.user?.role_name;
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const academicPeriodId = searchParams.period || "";
  const studyProgram = searchParams.studyProgram || "";
  const subjectId = params.subjectId;

  if (role !== DOSEN) notFound();

  const data = await getAllPresenceClassBySubjectId({
    academicPeriodId: academicPeriodId as string,
    subjectId,
    queryParam: {
      page: page,
      search: search as string,
    },
  });

  const presenceComponent = await getAllPresenceComponent({
    academicPeriodId: academicPeriodId as string,
    subjectId,
    studyProgramId: studyProgram as string,
  });

  if (!academicPeriodId || !studyProgram) redirect("/settings/presence");

  if (data?.status === 404) notFound();

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Pengaturan Presensi
            </h1>
            <TablePresenceClass
              data={data}
              presenceComponent={presenceComponent}
            />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
