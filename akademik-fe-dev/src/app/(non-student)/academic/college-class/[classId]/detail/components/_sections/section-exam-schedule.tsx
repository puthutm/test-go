"use client";
import React from "react";
// import third party component
import { Card, Button, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import DataTables from "@/components/ui/datatable";
import { FileDownloadIcon } from "@/components/icons/file-download";

import useColumnExamSchedule from "../_columns/column-definition-exam-schedule";

export interface IDummyValueTable {
  id: string;
  jenis_ujian: string;
  kelompok: string;
  tanggal: string;
  waktu: string;
  ruang: string;
  peserta: string;
}

const dummyValueTable: PaginationData<IDummyValueTable> = {
  metadata: {
    page: 1,
    size: 1,
    total_data: 1,
    total_page: 1,
  },
  data: [
    {
      id: "1",
      jenis_ujian: "text",
      kelompok: "text",
      tanggal: "text",
      waktu: "text",
      ruang: "text",
      peserta: "text",
    },
  ],
};
function SectionExamSchedule() {
  // define column for datatable
  const { columns } = useColumnExamSchedule();

  // dumy data

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Peserta Kelas
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
        <section className="table-responsive">
          <DataTables
            columns={columns}
            data={dummyValueTable}
            pageCount={0}
            pagination={null}
            setPagination={() => {}}
            isLoading={false}
            total={0}
            isPaginate={false}
          />
        </section>
      </Card>
    </section>
  );
}

export default SectionExamSchedule;
