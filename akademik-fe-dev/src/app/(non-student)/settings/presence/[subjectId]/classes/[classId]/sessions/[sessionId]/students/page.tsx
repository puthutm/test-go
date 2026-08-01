import { Card, Col, Row } from "reactstrap";
import { getServerSession } from "next-auth";
import { notFound } from "next/navigation";

import authOptions from "@/config/next-auth";
import { DOSEN } from "@/lib/constants/role";
import { getAllStudentPresenceBySessionId } from "@/services/api/settings/presence/lecturer/get-all-students-presence-by-session-id";
import { TableStudentPresenceBySessionId } from "./components/table-student-presence-by-session-id";

export default async function StudentPresenceBySessionId({
  searchParams,
  params,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  params: { sessionId: string };
}) {
  const session = await getServerSession(authOptions);
  const role = session?.user?.role_name;
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const sessionId = params.sessionId;

  const data = await getAllStudentPresenceBySessionId(sessionId, {
    page: page,
    search: search as string,
  });

  if (data?.status === 404 || role !== DOSEN) notFound();

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Pengaturan Presensi
            </h1>
            <TableStudentPresenceBySessionId data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
