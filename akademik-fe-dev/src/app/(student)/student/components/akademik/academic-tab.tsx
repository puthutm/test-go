import { Card, CardBody, Col, Row } from "reactstrap";

import { SideTab } from "../side-tabs";
import { DescriptionIcon } from "@/components/icons/description";
import { AssignmentIcon } from "@/components/icons/assignment";
import { GroupIcon } from "@/components/icons/group";
import { SummarizeIcon } from "@/components/icons/summarize";
import { FormPembimbing } from "./form-pemibimbing-akademik";
import FormKrs from "./form-krs";
import TableKHS from "./table-khs";
import { EmptyStateConversionGrade } from "./empty-state-conversion-grade";
import { TaskIcon } from "@/components/icons/task";
import { TableKrs } from "./table-krs";

export default function TabAcademic({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const stateParam = searchParams?.state;

  const tabs = [
    {
      label: "Hasil Studi",
      key: "khs",
      icon: (
        <DescriptionIcon
          color={stateParam === "khs" || !stateParam ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Pengisian KRS",
      key: "filling-krs",
      icon: (
        <TaskIcon color={stateParam === "filling-krs" ? "white" : "#495057"} />
      ),
    },
    {
      label: "Kartu Rencana Studi",
      key: "krs",
      icon: (
        <AssignmentIcon color={stateParam === "krs" ? "white" : "#495057"} />
      ),
    },
    {
      label: "Pembimbing Akademik",
      key: "academic-counselors",
      icon: (
        <GroupIcon
          color={stateParam === "academic-counselors" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Nilai Konversi",
      key: "conversion-value",
      icon: (
        <SummarizeIcon
          color={stateParam === "conversion-value" ? "white" : "#495057"}
        />
      ),
    },
  ];

  return (
    <Row className="pt-4">
      <Col lg={3}>
        <Card style={{ position: "sticky", top: 80, borderRadius: "8px" }}>
          <CardBody>
            <SideTab tabs={tabs} />
          </CardBody>
        </Card>
      </Col>
      <Col lg={9}>
        <Card style={{ borderRadius: "8px" }}>
          <CardBody>
            {stateParam === "khs" || !stateParam ? <TableKHS /> : null}
            {stateParam === "krs" ? <TableKrs /> : null}
            {stateParam === "filling-krs" ? <FormKrs /> : null}
            {stateParam === "academic-counselors" ? <FormPembimbing /> : null}
            {stateParam === "conversion-value" ? (
              <EmptyStateConversionGrade />
            ) : null}
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}
