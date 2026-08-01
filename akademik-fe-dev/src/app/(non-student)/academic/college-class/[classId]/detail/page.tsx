// import third part component
import { Row, Col } from "reactstrap";

// import component
import DetailClassSection from "./components/_sections/section-detail-class";
import SectionTeachingLecturer from "./components/_sections/section-teaching-lecturer";
import SectionClassParticipants from "./components/_sections/section-class-participants";
import SectionCollegeContract from "./components/_sections/section-college-contract";
import SectionExamSchedule from "./components/_sections/section-exam-schedule";
import SectionPresenceClass from "./components/_sections/section-presence-class";
import SectionCourseGrades from "./components/_sections/section-course-grades";
import SectionQuestionnaireRecap from "./components/_sections/section-questionnaire-recap";
import SectionRPS from "./components/_sections/section-rps";
import SectionCourseWork from "./components/_sections/section-course-work";
import SectionClassSchedule from "./components/_sections/section-class-schedule";

function DetailCollegeClass({
  searchParams,
}: {
  params: {
    classId: string;
  };
  searchParams: {
    tab: string;
  };
}) {
  return (
    <Row>
      <Col>
        {searchParams.tab === "class" ? (
          <DetailClassSection />
        ) : searchParams.tab === "lecturer" ? (
          <SectionTeachingLecturer />
        ) : searchParams.tab === "member" ? (
          <SectionClassParticipants />
        ) : searchParams.tab === "contract" ? (
          <SectionCollegeContract />
        ) : searchParams.tab === "schedule-college" ? (
          <SectionClassSchedule />
        ) : searchParams.tab === "schedule-exam" ? (
          <SectionExamSchedule />
        ) : searchParams.tab === "presence" ? (
          <SectionPresenceClass />
        ) : searchParams.tab === "course-grades" ? (
          <SectionCourseGrades />
        ) : searchParams.tab === "questionnaire-recap" ? (
          <SectionQuestionnaireRecap />
        ) : searchParams.tab === "rps" ? (
          <SectionRPS />
        ) : searchParams.tab === "course-work" ? (
          <SectionCourseWork />
        ) : (
          <p>not found</p>
        )}
      </Col>
    </Row>
  );
}

export default DetailCollegeClass;
