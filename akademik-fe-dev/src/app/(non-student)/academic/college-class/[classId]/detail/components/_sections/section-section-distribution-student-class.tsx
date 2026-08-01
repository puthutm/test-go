"use client";
import React from "react";
import { useParams } from "next/navigation";
import { Row, Col, Spinner } from "reactstrap";

import { useGetStudentClassDistribution } from "@/services/api/academic/lecturer/class-schedule/detail-class/student-class-distribution/use-get-student-class-distribution";

function SectionDistributionStudentClass() {
  const params = useParams();
  const {
    data: dataStudentClassDistribution,
    isFetching: isLoadingStudentClassDistribution,
  } = useGetStudentClassDistribution(params.classId as string);
  return (
    <section className="position-relative">
      {/*//! title */}
      <h2 className="fs-5 fw-bold" style={{ color: "#3A3A3A" }}>
        Semua Kelas Mahasiswa
      </h2>

      {/*//! list */}
      <Row className="row-gap-3 my-4 px-2">
        {isLoadingStudentClassDistribution ? (
          <Spinner className="mx-auto" />
        ) : (
          dataStudentClassDistribution?.data?.map(
            (el: StudentClassDistribution) => {
              return (
                <Col key={el.id} sm={6}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-5"
                    style={{ color: "#495057" }}
                  >
                    {el.name ?? "-"}
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

export default SectionDistributionStudentClass;
