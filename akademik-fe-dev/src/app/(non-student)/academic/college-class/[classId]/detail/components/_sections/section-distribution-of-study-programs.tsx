"use client";
import React from "react";
import { useParams } from "next/navigation";
import { Row, Col, Spinner } from "reactstrap";

import { useGetDistributionOfStudyProgram } from "@/services/api/academic/lecturer/class-schedule/detail-class/distribution-of-study-programs/use-get-distribution-of-study-program";

function SectionDistributionOfStudyPrograms() {
  const params = useParams();
  const {
    data: dataDistributionOfStudent,
    isFetching: isLoadingDistributionOfStudent,
  } = useGetDistributionOfStudyProgram(params.classId as string);

  return (
    <section className="position-relative">
      {/*//! list */}
      <Row className="row-gap-3 my-4 px-2">
        {isLoadingDistributionOfStudent ? (
          <Spinner className="mx-auto" />
        ) : (
          dataDistributionOfStudent?.data?.map(
            (el: DistributionOfStudyProgram) => {
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

export default SectionDistributionOfStudyPrograms;
