"use client";
import React from "react";
import { useParams } from "next/navigation";
// import third party component
import { Card, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import DataTables from "@/components/ui/datatable";

import useColumnClassParticipants from "../_columns/column-definition-class-participants";

import { useGetClassParticipantDetailClass } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-participants/use-get-class-participant";

function SectionClassParticipants() {
  const params = useParams();
  const { data: dataClassParticipant, isLoading: isLoadingClassParticipant } =
    useGetClassParticipantDetailClass(params.classId as string);

  // column definition
  const { columns } = useColumnClassParticipants();

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0  border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Peserta Kelas
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

      {/*//! TABLES */}
      <Card className="p-4 rounded-3 bg-white">
        <section className="table-responsive">
          <DataTables
            columns={columns}
            data={dataClassParticipant?.data}
            isPaginate={false}
            pageCount={0}
            pagination={0}
            setPagination={() => {
              return;
            }}
            isLoading={isLoadingClassParticipant}
            total={0}
          />
        </section>
      </Card>
    </section>
  );
}

export default SectionClassParticipants;
