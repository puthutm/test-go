"use client";
import React from "react";
// import third party component
import { Card, CardBody, CardHeader, Alert, Col, Row } from "reactstrap";
import Link from "next/link";

// import component
import DataTables from "@/components/ui/datatable";

import useColumnDetailPresenceClass from "./components/column-definition-detail-presence-class";

export interface IDummyValueTable {
  id: string;
  nim: string;
  nama: string;
  keterangan: string;
  status: string;
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
      nim: "200101072087",
      nama: " Agus Putri Dayanti Zalukhu ",
      keterangan: "Manual Presensi EdLink",
      status: "Hadir",
    },
  ],
};
function PageDetailPresenceClass() {
  // columns
  const { columns } = useColumnDetailPresenceClass();
  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 pb-3 border-2 ">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Detail Presensi
            </h2>
          </section>

          {/*//! alert panduan pengisian nilai*/}
          <Alert
            color="secondary"
            className="w-100 py-2 px-3  m-0 mt-2"
            fade={false}
          >
            {/*line */}
            <div
              style={{
                position: "absolute",
                width: "2px",
                left: "0",
                top: "0",
                bottom: "0",
                background: "#489CF0",
              }}
            />
            <Row className="gap-2">
              <Col>
                <h2 className="m-0 p-0  fs-6" style={{ color: "#489CF0" }}>
                  Panduan 200201204 - Pengantar Teknologi Sistem Informasi
                  (SI102)
                </h2>
              </Col>
              <Col>
                <h2 className="m-0 p-0  fs-6" style={{ color: "#489CF0" }}>
                  Rabu, 23 Oktober 2024, 10:00 s.d 11:40
                </h2>
              </Col>
            </Row>

            {/* list */}
            <p
              className="m-0 p-0 mt-2 tetxt-underline"
              style={{ color: "#489CF0" }}
            >
              Saat ini Anda tidak dapat melakukan presensi mahasiswa karena
              nilai sudah dikunci. Silakan buka kunci nilai kelas terlebih
              dahulu{" "}
              <Link className="fw-bold" href={"/"}>
                klik disini
              </Link>
              .
            </p>
          </Alert>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
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
        </CardBody>
      </Card>
    </section>
  );
}

export default PageDetailPresenceClass;
