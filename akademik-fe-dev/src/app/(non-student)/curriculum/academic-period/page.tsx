import { Card, Col, Row } from "reactstrap";
import { Metadata } from "next";
import { TableAcademicPeriod } from "./components/table-academic-period";
import { getAcademicPeriods } from "@/services/api/data-referensi/academic-period/get-all-academic-period";

export const metadata: Metadata = {
  title: "Periode Akademik",
};

export default async function AcademicPeriodPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const data = await getAcademicPeriods({
    page: page,
    filter: search as string,
  });

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <h1 className="fs-4 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
              Periode Akademik
            </h1>
            <TableAcademicPeriod data={data} />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
