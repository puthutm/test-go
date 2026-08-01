import React from "react";
import { Col, Row } from "reactstrap";

export default function AcademicPeriodDetailLayout({
  children,
}: {
  // params: Promise<{ academicPeriodId: string }>;
  children: React.ReactNode;
}) {
  return (
    <Row>
      <Col>{children}</Col>
    </Row>
  );
}
