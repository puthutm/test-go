import { getClassLecturerByClassId } from "@/services/api/settings/academic-period/class-lecturer/get-class-lecturer";
import React from "react";
import { Col, Row } from "reactstrap";
import { ModalClassLecturer } from "./modal-class-lecturer";

export default async function ClassLecturer({
  params,
  isDetail,
}: {
  params: Promise<{ academicPeriodId: string; classId: string }>;
  isDetail?: boolean;
}) {
  const classId = (await params).classId;
  const data = await getClassLecturerByClassId(classId);

  return (
    <div className="d-flex flex-column gap-2">
      <div className="d-flex justify-content-between align-items-center">
        <h2 className="fs-5 fw-semibold mb-1" style={{ color: "#3A3A3A" }}>
          Pengajar Kelas
        </h2>
        {!isDetail ? (
          <ModalClassLecturer data={data} classId={classId} />
        ) : null}
      </div>
      <Row>
        <Col sm={2}>Dosen Pengajar</Col>
        <Col>: {data?.data?.LecturerName || ""}</Col>
      </Row>
      <Row>
        <Col sm={2}>Dosen Pengganti</Col>
        <Col>: {data?.data?.SubtituteLecturerName || ""}</Col>
      </Row>
    </div>
  );
}
