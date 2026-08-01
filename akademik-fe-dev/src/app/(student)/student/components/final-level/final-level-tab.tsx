import { Card, CardBody, Col, Row } from "reactstrap";

import { SideTab } from "../side-tabs";
import { DescriptionIcon } from "@/components/icons/description";
import { FastCheckIcon } from "@/components/icons/fast_check";
import { ChatIcon } from "@/components/icons/chat";
import { AssignmentTurnedInIcon } from "@/components/icons/assignment-turned-in";
import { SchoolIcon } from "@/components/icons/school";
import ProposalFinalLevelView from "./views/proposal-view";
import ProgressFinalTaskView from "./views/progres-final-level-view";
import CounselingView from "./views/counseling-view";
import JudiciaryView from "./views/judiciary-view";
import GraduationView from "./views/graduation-view";

export default function TabFinalLevel({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const stateParam = searchParams?.state;

  const tabs = [
    {
      label: "Proposal TA",
      key: "proposal",
      icon: (
        <DescriptionIcon
          color={stateParam === "proposal" || !stateParam ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Progres TA",
      key: "ta",
      icon: <FastCheckIcon color={stateParam === "ta" ? "white" : "#495057"} />,
    },
    {
      label: "Bimbingan TA",
      key: "counseling",
      icon: (
        <ChatIcon color={stateParam === "counseling" ? "white" : "#495057"} />
      ),
    },
    {
      label: "Yudisium",
      key: "judiciary",
      icon: (
        <AssignmentTurnedInIcon
          color={stateParam === "judiciary" ? "white" : "#495057"}
        />
      ),
    },
    {
      label: "Wisuda",
      key: "graduation",
      icon: (
        <SchoolIcon color={stateParam === "graduation" ? "white" : "#495057"} />
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
            {stateParam === "proposal" || !stateParam ? (
              <ProposalFinalLevelView />
            ) : null}
            {stateParam === "ta" ? <ProgressFinalTaskView /> : null}
            {stateParam === "counseling" ? <CounselingView /> : null}
            {stateParam === "judiciary" ? <JudiciaryView /> : null}
            {stateParam === "graduation" ? <GraduationView /> : null}
          </CardBody>
        </Card>
      </Col>
    </Row>
  );
}
