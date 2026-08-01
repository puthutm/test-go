import { Card, Col, Row } from "reactstrap";
import { Metadata } from "next";
import { getServerSession } from "next-auth";

import { TableLecturers } from "./components/table-lecturers";
import { getAllLecturers } from "@/services/api/program-head/portal/lecturers/get-all-lecturers";
import authOptions from "@/config/next-auth";
import { KAPRODI } from "@/lib/constants/role";

export const metadata: Metadata = {
  title: "Portal Dosen",
};

export default async function PortalStudentPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const session = await getServerSession(authOptions);
  const roleName = session?.user?.role_name;
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  let data;

  if (roleName === KAPRODI)
    data = await getAllLecturers({
      page: page,
      search: search as string,
    });

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              List Data Dosen
            </h1>
            <TableLecturers data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
