"use client";
import React, { useState } from "react";

// import third party component
import { Card, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import SectionWeeklySchedule from "./section-weekly-schedule";
import SectionDistributionOfStudyPrograms from "./section-distribution-of-study-programs";
import SectionDistributionOfLectureSystem from "./section-distribution-of-the-lecture-system";
import SectionDistributionStudentClass from "./section-section-distribution-student-class";
import InformationCollegeClass from "../information-college-class";

interface ITabs {
  id: string;
  title: string;
}

const listTabs: ITabs[] = [
  {
    id: "1",
    title: "Jadwal Mingguan",
  },
  // {
  //   id: "2",
  //   title: "Sebaran Program Studi",
  // },
  // {
  //   id: "3",
  //   title: "Sebaran Sistem Kuliah",
  // },
  // {
  //   id: "4",
  //   title: "Sebaran Kelas Mahasiswa",
  // },
];

function DetailClassSection() {
  const [activeTabs, setActiveTabs] = useState<string>("1");

  return (
    <section className="position-relative">
      {/*//! card detail class */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0  border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1">
            {/* //! text */}
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Data Kelas Kuliah
            </h2>
          </section>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
          {/*//!tabs */}
          <section
            className="position-relative mt-2"
            style={{ borderBottom: "2px solid #ddd" }}
          >
            <TabsSectionDetailCollegeClass />
          </section>

          {/*//! section informational */}
          <section
            className="position-relative mt-3"
            style={{ background: "#FAFCFF" }}
          >
            <InformationCollegeClass />
          </section>
        </CardBody>
      </Card>

      {/*//! card detail information*/}
      <Card className="py-3 px-4 mt-4 rounded-3 bg-white">
        {/*//! tabs */}
        <section className="position-relative mb-3 d-flex flex-wrap">
          {listTabs.map((tab: ITabs) => {
            return (
              <button
                key={tab.id}
                className={`btn-tabs-lecturer-subject  fw-medium
                        ${
                          tab.id === activeTabs
                            ? "btn-tabs-lecturer-subject-active"
                            : "btn-tabs-lecturer-subject-not-active"
                        }    
                            `}
                onClick={() => {
                  setActiveTabs(tab.id);
                }}
              >
                {tab.title}
              </button>
            );
          })}
        </section>

        {/*//! content */}
        <section className="position-relative mt-1">
          {activeTabs === "1" ? (
            <SectionWeeklySchedule />
          ) : activeTabs === "2" ? (
            <SectionDistributionOfStudyPrograms />
          ) : activeTabs === "3" ? (
            <SectionDistributionOfLectureSystem />
          ) : (
            <SectionDistributionStudentClass />
          )}
        </section>
      </Card>
    </section>
  );
}

export default DetailClassSection;
