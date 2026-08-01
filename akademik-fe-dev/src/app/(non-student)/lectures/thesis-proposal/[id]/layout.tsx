import React from "react";
import { Card, CardBody } from "reactstrap";
import ThesisProposalSidetab from "./_components/sidetab";

export default function ThesisProposalLayout({
  children,
  studentInfo,
}: {
  children: React.ReactNode;
  studentInfo: React.ReactNode;
}) {
  return (
    <div>
      <Card>
        <CardBody>
          <h4 className="fw-semibold">Proposal Tugas Akhir</h4>
          <div className="bg-info-subtle p-3 rounded mt-3">{studentInfo}</div>
        </CardBody>
      </Card>

      <div className="d-flex gap-3">
        <div className="position-relative w-25">
          <div className="bg-white p-3 rounded ">
          <ThesisProposalSidetab />

          </div>
        </div>
        <div className="bg-white p-3 rounded w-75">{children}</div>
      </div>
    </div>
  );
}
