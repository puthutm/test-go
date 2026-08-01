"use client";
import React from "react";
// import third party component
import { Card, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";

function SectionTeachingLecturer() {
  return (
    <section className="position-relative">
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0  border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1">
            {/* //! text */}
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Dosen Pengajar
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

          {/*//! informational college class */}
          <section className=" mt-3">
            <InformationCollegeClass />
          </section>
        </CardBody>
      </Card>
    </section>
  );
}

export default SectionTeachingLecturer;
