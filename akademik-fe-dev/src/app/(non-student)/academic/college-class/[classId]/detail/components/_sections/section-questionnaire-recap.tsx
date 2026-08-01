"use client";
import React from "react";
// import third party component
import { Card, CardBody, CardHeader } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import DataTables from "@/components/ui/datatable";

// import hook
import useColumnTotalRespondent from "../_columns/column-definition-total-respondent";

export interface IDummyValueTable {
  id: string;
  soal: string;
}

const dummyValueTable: PaginationData<IDummyValueTable> = {
  metadata: {
    page: 1,
    size: 1,
    total_data: 1,
    total_page: 1,
  },
  data: [],
};

function SectionQuestionnaireRecap() {
  //! column total respondent
  const { columns: columnTotalRespondent } = useColumnTotalRespondent();

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Rekap Kuesioner
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
        {/*//! total responden */}
        <section className="position-relative">
          <h2 className="fs-6 fw-bold" style={{ color: "#3A3A3A" }}>
            Jumlah Responden : 0/43
          </h2>
          <div className="table-responsive mt-3">
            <DataTables
              columns={columnTotalRespondent}
              data={dummyValueTable}
              pageCount={0}
              pagination={null}
              setPagination={() => {}}
              isLoading={false}
              total={0}
              isPaginate={false}
            />
          </div>
        </section>

        {/*//! open question */}
        <section className="position-relative">
          <h2 className="fs-6 fw-bold" style={{ color: "#3A3A3A" }}>
            Open Question
          </h2>
          <div className="table-responsive mt-3">
            <DataTables
              columns={columnTotalRespondent}
              data={dummyValueTable}
              pageCount={0}
              pagination={null}
              setPagination={() => {}}
              isLoading={false}
              total={0}
              isPaginate={false}
            />
          </div>
        </section>
      </Card>
    </section>
  );
}

export default SectionQuestionnaireRecap;
