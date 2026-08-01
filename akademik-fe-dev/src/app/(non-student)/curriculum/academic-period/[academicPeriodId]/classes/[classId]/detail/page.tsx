import { Card, CardBody, Col, Row } from "reactstrap";
import { notFound } from "next/navigation";

import { SideTab } from "../components/sidetab";
import ClassDetail from "../components/class-detail";
import ClassParticipant from "../components/class-participant";
import AcademicPeriodInfo from "@/components/ui/academic-period-info";
import ClassContract from "../components/class-contract";
import ClassLecturer from "../components/class-lecturer";
import { getDetailClassForProgramHead } from "@/services/api/curriculum/academic-period/class/get-detail-class";
import ClassSchedule from "../components/class-schedule";

export default async function DetailClassPage({
  params,
  searchParams,
}: {
  params: Promise<{ academicPeriodId: string; classId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const tabSearchParam = searchParams.tabs;

  const classId = (await params).classId;

  const detailClass = await getDetailClassForProgramHead(classId);

  if (detailClass.status === 404) return notFound();

  return (
    <div className="flex flex-column gap-2">
      <Row>
        <Col>
          <Card>
            <div className="d-flex flex-column gap-3 p-3">
              <AcademicPeriodInfo params={params} />
            </div>
          </Card>
        </Col>
      </Row>
      <Row>
        <Col lg={3}>
          <Card style={{ position: "sticky", top: 80, borderRadius: "8px" }}>
            <CardBody>
              <SideTab />
            </CardBody>
          </Card>
        </Col>
        <Col lg={9}>
          <Card style={{ borderRadius: "8px" }}>
            <CardBody>
              {tabSearchParam === "class-detail" || !tabSearchParam ? (
                <ClassDetail
                  params={params}
                  detailClass={detailClass}
                  isDetail
                />
              ) : null}
              {tabSearchParam === "participant" ? (
                <ClassParticipant
                  detailClass={detailClass}
                  searchParams={searchParams}
                  isDetail
                />
              ) : null}
              {tabSearchParam === "class-contract" ? (
                <ClassContract
                  params={params}
                  detailClass={detailClass}
                  isDetail
                />
              ) : null}
              {tabSearchParam === "class-lecturer" ? (
                <ClassLecturer params={params} isDetail />
              ) : null}
              {tabSearchParam === "class-schedule" ? (
                <ClassSchedule
                  classId={classId}
                  searchParams={searchParams}
                  isDetail
                />
              ) : null}
            </CardBody>
          </Card>
        </Col>
      </Row>
    </div>
  );
}
