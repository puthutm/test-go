import { Card, Col, Row } from "reactstrap";
import { Metadata } from "next";
import { getServerSession } from "next-auth";

import { TableStudent } from "./components/table-students";
import { getAllStudentsForProgramHead } from "@/services/api/program-head/portal/students/get-all-students";
import authOptions from "@/config/next-auth";
import { KAPRODI, MAHASISWA } from "@/lib/constants/role";
import { getAllStudentsForAcademic } from "@/services/api/portal/academic/get-all-student";
import { notFound } from "next/navigation";

export const metadata: Metadata = {
  title: "Portal Mahasiswa",
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

  if (roleName === MAHASISWA) notFound();

  if (roleName === KAPRODI) {
    data = await getAllStudentsForProgramHead({
      page: page,
      search: search as string,
    });
  } else {
    data = await getAllStudentsForAcademic({
      page: page,
      search: search as string,
    });
  }

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              List Data Mahasiswa
            </h1>
            <TableStudent data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
