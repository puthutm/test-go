"use client";
import React, { useState } from "react";

// import third party component
import { Row, Col } from "reactstrap";

// import component
// import InformationSubject from "../information-subject";
import DataTables from "@/components/ui/datatable";

import useColumnCPL from "../_columns/column-definition-cpl";
import Link from "next/link";

export interface IDummyValueTable {
  id: string;
  code: string;
  description: string;
}

export interface IModalManipulationFilterSubjects {
  status: boolean;
  title: "Filter";
  data?: string;
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
      code: "IF3110",
      description: "test",
    },
    {
      id: "2",
      code: "IF3110",
      description: "test",
    },
  ],
};

function SectionCPL() {
  const [queryParams] = useState<QueryParam>({
    page: 1,
  });

  //! column cpl
  const { columns } = useColumnCPL();

  return (
    <section className="position-relative mt-2">
      <Row className="row-gap-3">
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                CPL-Prodi
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                text here
              </p>
            </Col>
          </Row>
        </Col>
      </Row>

      {/*//! list */}
      <section className="position-relative mt-3">
        <h2 className="m-0 p-0 fw-semibold fs-6" style={{ color: "#3A3A3A" }}>
          CPMK
        </h2>
        <section className="my-2">
          <ol className="px-4 fw-medium" style={{ color: "#495057" }}>
            <li className="mb-1">
              Mempelajari Metode Pengujian Sistem Informasi: Tujuan utama dari
              mata kuliah ini adalah untuk mengajarkan mahasiswa tentang
              berbagai metode pengujian yang digunakan dalam menguji sistem
              informasi. Ini mencakup pengenalan terhadap konsep-konsep dasar
              dalam pengujian perangkat lunak, teknik pengujian, dan strategi
              pengujian.
            </li>

            <li className="mb-1">
              Perencanaan Implementasi Sistem Informasi: Mata kuliah ini juga
              bertujuan untuk memberikan pemahaman tentang bagaimana
              merencanakan dan melaksanakan implementasi sistem informasi di
              berbagai organisasi. Ini mencakup tahapan perencanaan, desain,
              pengembangan, dan pelaksanaan sistem informasi.
            </li>

            <li className="mb-1">
              Metode Pengujian Usabilitas: Mahasiswa akan mempelajari cara
              menguji dan mengevaluasi usabilitas dari sistem informasi. Hal ini
              melibatkan pemahaman tentang bagaimana pengguna berinteraksi
              dengan sistem, dan bagaimana membuat sistem lebih mudah digunakan
              dan efisien.
            </li>

            <li className="mb-1">
              Laporan Perencanaan, Kontrol, dan Evaluasi Pengujian Sistem
              Informasi: Tujuan lainnya adalah mengajarkan cara merancang
              laporan perencanaan, kontrol, dan evaluasi dari pengujian sistem
              informasi. Ini termasuk pemantauan progres pengujian, mengukur
              keefektifan pengujian, dan menyusun laporan hasil pengujian.
            </li>
          </ol>

          <p className="m-0 p-0 fw-medium" style={{ color: "#495057" }}>
            Mata kuliah ini bertujuan untuk mempersiapkan mahasiswa dengan
            pengetahuan dan keterampilan yang diperlukan untuk merancang,
            menguji, dan mengimplementasikan sistem informasi dengan baik dan
            efisien, serta memastikan bahwa sistem tersebut memenuhi kebutuhan
            pengguna dan organisasi.
          </p>
        </section>
      </section>

      {/*//! tables cpl */}
      <section className="mt-3 position-relative table-responsive">
        <DataTables
          columns={columns}
          data={dummyValueTable}
          pageCount={dummyValueTable?.metadata.total_page as number}
          pagination={queryParams}
          setPagination={() => {}}
          isLoading={false}
          total={dummyValueTable.metadata.total_data as number}
        />
      </section>

      {/*//! BUTTON BACK  */}
      <section className="position-relative mt-3 d-flex justify-content-starts">
        <Link
          href={"?tab=subject"}
          className="btn  d-flex align-items-center gap-2"
          color="transparent"
          style={{ color: "#10487A", border: "1px solid #10487A" }}
        >
          Kembali
        </Link>
      </section>
    </section>
  );
}

export default SectionCPL;
