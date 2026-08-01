"use client";
import React from "react";
import { useParams } from "next/navigation";
import { Row, Col, Spinner } from "reactstrap";

import { useGetAcademicSystemDistribution } from "@/services/api/academic/lecturer/class-schedule/detail-class/academic-system-distribution/use-get-academic-system-distribution";

function SectionDistributionOfLectureSystem() {
  const params = useParams();
  const {
    data: dataAcademicSystemDistribution,
    isFetching: isLoadingAcademicSystemDistribution,
  } = useGetAcademicSystemDistribution(params.classId as string);

  return (
    <section className="position-relative">
      {/*//! title */}
      <h2 className="fs-5 fw-bold" style={{ color: "#3A3A3A" }}>
        Semua Sistem Kuliah
      </h2>

      {/*//! list */}
      <Row className="row-gap-3 my-4 px-2">
        {isLoadingAcademicSystemDistribution ? (
          <Spinner className="mx-auto" />
        ) : (
          dataAcademicSystemDistribution?.data?.map(
            (el: AcademicSystemDistribution, index: number) => {
              return (
                <Col key={index} sm={6}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-5"
                    style={{ color: "#495057" }}
                  >
                    {el.type_of_meeting}
                  </h3>
                </Col>
              );
            }
          )
        )}
      </Row>
    </section>
  );
}

export default SectionDistributionOfLectureSystem;
