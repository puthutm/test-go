import { Card, CardBody, Col, Row } from "reactstrap";
import { SideTab } from "../side-tabs";
import { FolderIcon } from "@/components/icons/folder";
import { TableLetter } from "./table-letter";
import { DescriptionIcon } from "@/components/icons/description";

export default function TabAdministration({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const stateParam = searchParams?.state;
  const tabs = [
    {
      label: "Persuratan",
      key: "mail",
      icon: (
        <FolderIcon
          color={stateParam === "mail" || !stateParam ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Prasyarat Perkuliahan",
      key: "requirement",
      icon: (
        <DescriptionIcon
          color={stateParam === "requirement" ? "white" : "#495057"}
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
            {stateParam === "mail" || !stateParam ? <TableLetter /> : null}
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}
