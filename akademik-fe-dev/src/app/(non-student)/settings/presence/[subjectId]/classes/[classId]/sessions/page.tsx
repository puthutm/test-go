import React from "react";
import { Card, Col, Row } from "reactstrap";
import { getServerSession } from "next-auth";
import { notFound, redirect } from "next/navigation";

import authOptions from "@/config/next-auth";
import { DOSEN } from "@/lib/constants/role";
import { getAllPresenceClassSession } from "@/services/api/settings/presence/lecturer/get-all-presence-class-session";
import { TablePresenceClassSession } from "./components/table-presence-class-session";

export default async function PresenceClassSessionPage({
  searchParams,
  params,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  params: { subjectId: string; classId: string };
}) {
  const session = await getServerSession(authOptions);
  const role = session?.user?.role_name;
  const academicPeriodId = searchParams.period || "";
  const studyProgram = searchParams.studyProgram || "";
  const subjectId = params.subjectId;
  const classId = params.classId;

  if (role !== DOSEN) notFound();

  const data = await getAllPresenceClassSession({
    academicPeriodId: academicPeriodId as string,
    subjectId,
    classId,
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
            <TablePresenceClassSession data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
