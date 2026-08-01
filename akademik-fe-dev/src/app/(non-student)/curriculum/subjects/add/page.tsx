import React from "react";
import { Card, CardBody, CardHeader } from "reactstrap";
import SubjectForm from "./_component/SubjectForm";

export default function AddSubjectPage() {
  return (
    <Card>
      <CardHeader>
        <h4 className="fw-bold">Tambah Mata Kuliah</h4>
      </CardHeader>
      <CardBody>
        <SubjectForm />
      </CardBody>
    </Card>
  );
}
