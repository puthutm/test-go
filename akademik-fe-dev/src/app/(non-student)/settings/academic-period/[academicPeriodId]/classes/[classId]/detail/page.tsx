import { Card, CardBody, Col, Row } from "reactstrap";
import { notFound } from "next/navigation";
import { Metadata } from "next";

import ClassDetail from "../components/class-detail";
import AcademicPeriodInfo from "@/components/ui/academic-period-info";
import { SideTab } from "../components/sidetab";
import ClassParticipant from "../components/class-participant";
import ClassContract from "../components/class-contract";
import { getDetailClass } from "@/services/api/settings/academic-period/class/get-detail-class";
import ClassLecturer from "../components/class-lecturer";
import ClassSchedule from "../components/class-schedule";
import ClassScoreView from "../components/class-score";

export const metadata: Metadata = {
  title: "Detail Kelas",
  description: "Detail Kelas",
};
export default async function DetailSubjectPage({
  params,
  searchParams,
}: {
  params: Promise<{ academicPeriodId: string; classId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const tabSearchParam = searchParams.tabs;

  const classId = (await params).classId;

  const detailClass = await getDetailClass(classId);

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
              {tabSearchParam === "class-score" && (
                <ClassScoreView
                  detailClass={detailClass}
                  searchParams={searchParams}
                />
              )}
            </CardBody>
          </Card>
        </Col>
      </Row>
    </div>
  );
}
