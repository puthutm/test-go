"use client";
import React from "react";
// import third party component
import {
  Card,
  CardBody,
  CardHeader,
  Button,
  Label,
  Row,
  Col,
} from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";

interface IList {
  id: number;
  title: string;
}

const listIntruksiUmum: IList[] = [
  {
    id: 1,
    title: "Mahasiswa mampu merancang pemrograman",
  },
  {
    id: 2,
    title: "Menggunakan visual basic dengan baik",
  },
  {
    id: 3,
    title: "Mampu menggunakan model perancangan yang benar",
  },
  {
    id: 4,
    title:
      "Berdasarkan konsep-konsep pemrograman visual menggunakan visual basic.",
  },
];

const listBahanAjar: IList[] = [
  {
    id: 1,
    title:
      "Konsep Dasar Pemrograman Visual: pengertian dan tujuan dari pemrograman  visual, konsep dasar pengembangan aplikasi berbasis event, pemrograman  berorientasi objek, dan paradigma pemrograman visual.",
  },
  {
    id: 2,
    title:
      "Pengenalan Bahasa Pemrograman Visual: pemahaman tentang bahasa pemrograman visual, seperti C#, Visual Basic, JavaFX, dan Python.",
  },
  {
    id: 3,
    title:
      "Teknik Pengembangan Antarmuka Grafis: desain antarmuka pengguna, komponen antarmuka, layout, dan pengaturan tampilan aplikasi.",
  },
  {
    id: 4,
    title:
      "Manipulasi Gambar: manipulasi gambar digital, seperti cropping, resizing, rotating, dan flipping.",
  },
  {
    id: 5,
    title:
      "Pengolahan Multimedia: integrasi multimedia, seperti audio, video, dan animasi dalam aplikasi",
  },
  {
    id: 6,
    title:
      "Pengembangan Aplikasi Desktop Interaktif: pengembangan aplikasi desktop  interaktif dengan menggunakan teknologi visual dan multimedia, termasuk  integrasi dengan database.",
  },
  {
    id: 7,
    title:
      "Implementasi dan Evaluasi Aplikasi: proses implementasi aplikasi, pengujian, debugging, dan evaluasi aplikasi.",
  },
];

function SectionRPS() {
  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">RPS</h2>
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

      {/*//! content */}
      <Card className="p-4 rounded-3 bg-white">
        {/* list instruksional umum*/}
        <div className="position-relative">
          <h2
            className="m-0 p-0 fw-semibold "
            style={{ fontSize: "13px", color: "#495057" }}
          >
            Tujuan Instruksional Umum
          </h2>
          <ol className="mt-2 px-4">
            {listIntruksiUmum?.map((el: IList) => {
              return (
                <li
                  key={el.id}
                  className=""
                  style={{ fontSize: "13px", color: "#495057" }}
                >
                  {el.title}
                </li>
              );
            })}
          </ol>
        </div>

        {/* list bahan ajar*/}
        <div className="position-relative">
          <h2
            className="m-0 p-0 fw-semibold "
            style={{ fontSize: "13px", color: "#495057" }}
          >
            Bahan Ajar
          </h2>
          <p
            className="m-0 p-0 "
            style={{ fontSize: "13px", color: "#495057" }}
          >
            Berikut ini adalah beberapa materi pembelajaran yang akan dibahas
            dalam mata kuliah Pemrograman Visual.
          </p>
          <ol className="px-4">
            {listBahanAjar?.map((el: IList) => {
              return (
                <li
                  key={el.id}
                  className=""
                  style={{ fontSize: "13px", color: "#495057" }}
                >
                  {el.title}
                </li>
              );
            })}
          </ol>

          <p
            className="m-0 p-0 mt-3"
            style={{ fontSize: "13px", color: "#495057" }}
          >
            Materi pembelajaran ini akan membantu mahasiswa memahami cara
            mengembangkan aplikasi desktop yang menarik, interaktif, dan mudah
            digunakan dengan menggunakan teknologi visual dan multimedia.
            Mahasiswa juga akan mempelajari teknik-teknik pemrograman dan desain
            antarmuka yang akan membantu mereka menjadi pengembang aplikasi yang
            lebih baik.
          </p>
        </div>

        {/* button dokumen */}
        <Row className="row-gap-1 mt-3">
          <Col sm={12}>
            <h3
              className="m-0 p-0 fw-semibold fs-6"
              style={{ color: "#3A3A3A" }}
            >
              Dokument RPS
            </h3>
          </Col>
          <Col sm={12}>
            <div className="">
              <input
                // ref={inputImport}
                type="file"
                // onChange={handleImport}
                style={{ display: "none" }}
                name="fImport"
                id="Fimport"
                accept=".xlsx, .xls, .csv"
              />
              <Button
                type="button"
                // disabled={isFetchingExportUser || isPendingImportUser}
                className="btn create-btn d-flex align-items-center p-0"
                style={{ background: "#44A7FF" }}
                // onClick={handleShow}
                data-bs-target="#api-key-modal"
              >
                <Label
                  htmlFor="Fimport"
                  className="py-1 px-3 m-0 pointer fs-6"
                  style={{ cursor: "pointer" }}
                >
                  DokumenRPS.pdf
                </Label>
              </Button>
            </div>
          </Col>
        </Row>
      </Card>
    </section>
  );
}

export default SectionRPS;
