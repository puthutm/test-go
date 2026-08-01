import React from "react";
import { Card, Col, Row } from "reactstrap";
import { getServerSession } from "next-auth";
import { notFound } from "next/navigation";

import authOptions from "@/config/next-auth";
import { DOSEN } from "@/lib/constants/role";
import { TablePresence } from "./components/table-presence";
import { getAllPresenceSubject } from "@/services/api/settings/presence/lecturer/get-all-presence-subjects";

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

  if (role !== DOSEN) notFound();

  const data = await getAllPresenceSubject({
    academic_periode_id: academicPeriodId as string,
    page: page,
    search: search as string,
  });
  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Pengaturan Presensi
            </h1>
            <TablePresence data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
