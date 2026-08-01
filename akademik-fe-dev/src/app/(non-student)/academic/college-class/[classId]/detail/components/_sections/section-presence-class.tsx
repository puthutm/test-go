"use client";
import React from "react";
import { useParams } from "next/navigation";
// import third party component
import { Card, CardBody, CardHeader, Button } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import { FileDownloadIcon } from "@/components/icons/file-download";
import DataTables from "@/components/ui/datatable";

import useColumnPresenceClass from "../_columns/column-definition-presence-class";

import { useGetClassAttendance } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-attendance/use-get-class-attendance";

function SectionPresenceClass() {
  //! columns
  const { columns } = useColumnPresenceClass();
  const params = useParams();

  //! get data
  const { data: dataClassAttendance, isLoading: isLoadingClassAttendance } =
    useGetClassAttendance(params.classId as string);

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Presensi Kelas
            </h2>
            <Button
              className="btn btn-primary d-flex align-items-center gap-2"
              color="#10487A"
              style={{
                padding: "6px 10px",
              }}
            >
              <FileDownloadIcon color="#fff" />
              Download
            </Button>
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
        {/*//! datatable */}
        <section className="table-responsive mt-3">
          <DataTables
            columns={columns}
            data={dataClassAttendance?.data}
            pageCount={0}
            pagination={null}
            setPagination={() => {}}
            isLoading={isLoadingClassAttendance}
            total={0}
            isPaginate={false}
          />
        </section>
      </Card>
    </section>
  );
}

export default SectionPresenceClass;
