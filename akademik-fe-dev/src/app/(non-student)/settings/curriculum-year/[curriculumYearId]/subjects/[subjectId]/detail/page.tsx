import React from "react";
import { Card, CardBody, CardHeader } from "reactstrap";
import { notFound } from "next/navigation";

import { FormSubject } from "../../components/form-subject";
import { Tabs } from "./components/tabs";
import { getSubjectById } from "@/services/api/settings/subject/get-subject-by-id";

export default async function EditSubject({
  params,
  searchParams,
}: {
  params: Promise<{ subjectId: string }>;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const subjectId = (await params).subjectId;
  const tabParams = searchParams.tabs as string;
  const detailSubject = await getSubjectById(subjectId);

  if (detailSubject?.status === 404) {
    return notFound();
  }

  return (
    <Card>
      <CardHeader className="p-0 mx-4 mt-4">
        <h1 className="fs-5" style={{ fontWeight: "500", color: "#3A3A3A" }}>
          Detail Mata Kuliah
        </h1>
      </CardHeader>
      <CardBody className="px-4">
        <Tabs param={subjectId} />
        {tabParams === "subject-data" || !tabParams ? (
          <FormSubject isDetail data={detailSubject?.data} />
        ) : null}
        {tabParams === "cpl-cpmk" ? <p>cpmk</p> : null}
        {tabParams === "rps" ? <p>RPS</p> : null}
      </CardBody>
    </Card>
  );
}
