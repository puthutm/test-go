import React from "react";
import { Card, CardBody, CardHeader } from "reactstrap";

import { Tabs } from "./components/tabs";
import { FormSubject } from "../components/form-subject";

export default function AddSubject({
  searchParams,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const tabParams = searchParams.tabs as string;

  return (
    <Card>
      <CardHeader className="p-0 mx-4 mt-4">
        <h1 className="fs-5" style={{ fontWeight: "500", color: "#3A3A3A" }}>
          Tambah Mata Kuliah
        </h1>
      </CardHeader>
      <CardBody className="px-4">
        <Tabs />
        {(tabParams === "subject-data" || !tabParams) && <FormSubject />}
        {tabParams === "cpl-cpmk" ? <p>cpmk</p> : null}
        {tabParams === "rps" ? <p>RPS</p> : null}
      </CardBody>
    </Card>
  );
}
