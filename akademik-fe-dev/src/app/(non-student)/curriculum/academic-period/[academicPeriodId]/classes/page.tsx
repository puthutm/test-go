import { Card, Col, Row } from "reactstrap";
import { TableSubjectAcademicPeriod } from "./components/table-subject-academic-period";

import AcademicPeriodInfo from "@/components/ui/academic-period-info";
import { getAllClassByAcademicPeriodIdForProgramHead } from "@/services/api/curriculum/academic-period/class/get-all-class";

export default async function DetailAcademicPeriodPage({
  params,
  searchParams,
}: {
  params: Promise<{ academicPeriodId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const academicPeriodId = (await params).academicPeriodId;

  const data = await getAllClassByAcademicPeriodIdForProgramHead({
    academic_periode_id: academicPeriodId as string,
    page: page,
    search: search as string,
  });

  return (
    <Row>
      <Col>
        <Card>
          <div className="d-flex flex-column gap-3 p-3">
            <AcademicPeriodInfo params={params} />
            <TableSubjectAcademicPeriod
              searchParams={searchParams}
              data={data}
            />
          </div>
        </Card>
      </Col>
    </Row>
  );
}
