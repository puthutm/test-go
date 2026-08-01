import { ChatIcon } from "@/components/icons/chat";
import { DescriptionIcon } from "@/components/icons/description";
import { InfoIcon } from "@/components/icons/info";
import { InsertDriveFileIcon } from "@/components/icons/inser-drive-file";
import { Card, CardBody, Col, Row } from "reactstrap";
import { SideTab } from "../side-tabs";
import InfoKpView from "./views/info-kp-view";
import { FormSeminar } from "./form-seminar-kp";
import CounselingKpViews from "./views/counseling-kp";
import LaporanKerjaView from "./views/laporan-kp-view";

export default function TabKp({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const stateParam = searchParams?.state;

  const tabs = [
    {
      label: "Informasi KP",
      key: "info",
      icon: (
        <InfoIcon
          color={stateParam === "info" || !stateParam ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Bimbingan KP",
      key: "counseling",
      icon: (
        <ChatIcon color={stateParam === "counseling" ? "white" : "#495057"} />
      ),
    },
    {
      label: "Seminar KP",
      key: "seminar",
      icon: (
        <InsertDriveFileIcon
          color={stateParam === "seminar" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Laporan KP",
      key: "report",
      icon: (
        <DescriptionIcon
          color={stateParam === "report" ? "white" : "#495057"}
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
            {stateParam === "info" || !stateParam ? <InfoKpView /> : null}
            {stateParam === "counseling" ? <CounselingKpViews /> : null}
            {stateParam === "seminar" ? <FormSeminar /> : null}
            {stateParam === "report" ? <LaporanKerjaView /> : null}
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}
