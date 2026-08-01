// import layout
import LayoutDetailSubject from "./components/_layout/layout-detail-subject";

// import compomnent
import { Card, CardBody } from "reactstrap";
import SectionSubject from "./components/_sections/section-subject";
import SectionCPL from "./components/_sections/section-cpl";
import SectionRPS from "./components/_sections/section-rps";
import SectionLessonPlan from "./components/_sections/section-lesson-plan";
import SectionEvaluationPlan from "./components/_sections/section-evaluation-plan";

function PageDetailSubject({
  searchParams,
}: {
  params: {
    id: string;
  };
  searchParams: {
    tab: string;
  };
}) {
  return (
    <LayoutDetailSubject>
      {searchParams.tab === "subject" ? (
        <SectionSubject />
      ) : searchParams.tab === "cpl" ? (
        <SectionCPL />
      ) : searchParams.tab === "rps" ? (
        <SectionRPS />
      ) : searchParams.tab === "study" ? (
        <SectionLessonPlan />
      ) : searchParams.tab === "evaluation" ? (
        <SectionEvaluationPlan />
      ) : (
        <Card>
          <CardBody className="d-flex justify-content-center align-items-center">
            <p>not found</p>
          </CardBody>
        </Card>
      )}
    </LayoutDetailSubject>
  );
}

export default PageDetailSubject;
