"use client";
import React from "react";
// import third party component
import { Card, CardBody, CardHeader, Button, Alert } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";
import { FileDownloadIcon } from "@/components/icons/file-download";
import DataTables from "@/components/ui/datatable";

import useColumnDefinitionCourseGrade from "../_columns/column-definition-course-grade";

export interface IDummyValueTable {
  id: string;
  nim: string;
  nama: string;
  hadir: string;
  tugas: string;
  uts: string;
  uas: string;
  kehadiran: string;
  nilai: string;
  grade: string;
  lulus: string;
  keterangan: string;
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
      nim: "text",
      nama: "text",
      hadir: "00.0",
      tugas: "text",
      uts: "text",
      uas: "text",
      kehadiran: "text",
      nilai: "text",
      grade: "text",
      lulus: "text",
      keterangan: "text",
    },
  ],
};

interface IPanduanPengisianNilai {
  id: number;
  title: string;
}

const panduanPengisianNilai: IPanduanPengisianNilai[] = [
  {
    id: 1,
    title:
      "Untuk mengisikan nilai silakan klik button edit atau mengupload nilai.",
  },
  {
    id: 2,
    title:
      "Untuk mengupload nilai, download template excel-nya terlebih dahulu pada button aksi",
  },
  {
    id: 3,
    title:
      "Pengisian nilai dilakukan pada tiap komposisi nilai sesuai dengan yang ditentukan tiap jenis mata kuliah pada tiap-tiap prodi dan kurikulum.",
  },
  {
    id: 4,
    title:
      "Untuk menyimpan data klik button simpan, pada proses ini nilai masih dapat diubah.",
  },
  {
    id: 5,
    title:
      "Untuk menampilkan nilai tiap komposisi nilai kepada mahasiswa klik umumkan nilai pada button aksi",
  },
  {
    id: 6,
    title:
      "Untuk menampilkan nilai akhir kepada mahasiswa klik kunci nilai pada button aksi.",
  },
  {
    id: 7,
    title:
      "Jika nilai sudah dikunci dosen tidak dapat melakukan pengubahan nilai",
  },
];
function SectionCourseGrades() {
  //! columns
  const { columns } = useColumnDefinitionCourseGrade();

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Nilai Perkuliahan
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
        {/*//! alert panduan pengisian nilai*/}
        <Alert color="secondary" className="w-100 py-2 px-3  m-0" fade={false}>
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
          <h2 className="m-0 p-0 fw-bold fs-6" style={{ color: "#489CF0" }}>
            Panduan Pengisian Nilai.
          </h2>

          {/* list */}
          <ol className="mt-2" style={{ color: "#489CF0" }}>
            {panduanPengisianNilai?.map((el: IPanduanPengisianNilai) => {
              return (
                <li key={el.id} className="fw-semibold">
                  {el.title}
                </li>
              );
            })}
          </ol>
        </Alert>

        {/*//! alert warning */}
        <Alert
          color="warning"
          className="w-100 py-2 px-3 d-flex gap-2 align-items-center m-0 mt-3"
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
              background: "#F7B84B",
            }}
          />
          <p className="m-0 p-0 flex-grow-1">
            Untuk mengembalikan komposisi nilai sesuai yang diatur di prodi
            silahkan klik di sini kemudian klik {`"Reset Komposisi"`}. Jika
            terdapat informasi {`"Sunting KRS"`} pada kolom keterangan, maka
            nilai akhir mahasiswa tidak akan terhitung ulang (perhitungan
            komposisi nilai dapat tidak sama dengan nilai akhir).
          </p>
        </Alert>

        {/*//! datatable */}
        <section className="table-responsive mt-3">
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

export default SectionCourseGrades;
