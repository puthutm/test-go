"use client";
import React from "react";

// import third party component
import { Row, Col, Button, Label } from "reactstrap";
import Link from "next/link";

import { PrintIcon } from "@/components/icons/print";

function SectionRPS() {
  return (
    <section className="position-relative mt-2">
      {/* button action */}
      <section className="d-flex justify-content-end">
        {/*//! action */}
        <section className="d-flex align-items-center gap-2">
          {/*//! print btn */}
          <Button
            className="btn  d-flex align-items-center gap-2 "
            color="transparent"
            style={{
              color: "#10487A",
              padding: "5.5px 12px",
              border: "1px solid #10487A",
            }}
          >
            <PrintIcon color="#10487A" />
            Print
          </Button>
        </section>
      </section>
      {/*//! tanggal penyusunan */}
      <Row className="gap-1">
        <Col sm={12}>
          <h3 className="m-0 p-0 fw-semibold fs-6" style={{ color: "#3A3A3A" }}>
            Tanggal Penyusunan
          </h3>
        </Col>
        <Col sm={12}>
          <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
            text here
          </p>
        </Col>
      </Row>

      {/*//! deskripsi mata kuliah ind */}
      <Row className="gap-1 mt-3">
        <Col sm={12}>
          <h3 className="m-0 p-0 fw-semibold fs-6" style={{ color: "#3A3A3A" }}>
            Deskripsi Mata Kuliah {`(IND)`}
          </h3>
        </Col>
        <Col sm={12}>
          <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
            Mata kuliah mempelajari metode pengujian sistem informasi,
            perencanaan implementasi sistem informasi, metode pengujian
            usabilitas dan laporan perencanaan, kontrol dan evaluasi pengujian
            sistem informasi
          </p>
        </Col>
      </Row>

      {/*//! deskripsi mata kuliah ind */}
      <Row className="gap-1 mt-3">
        <Col sm={12}>
          <h3 className="m-0 p-0 fw-semibold fs-6" style={{ color: "#3A3A3A" }}>
            Deskripsi Mata Kuliah {`(EG)`}
          </h3>
        </Col>
        <Col sm={12}>
          <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
            This course covers methods for information system testing, planning
            for information system implementation, usability testing methods,
            and reporting on planning, control, and evaluation of information
            system testing.
          </p>
        </Col>
      </Row>

      {/*//! list */}
      <section className="position-relative mt-3">
        <h2 className="m-0 p-0 fw-semibold fs-6" style={{ color: "#3A3A3A" }}>
          Tujuan Mata Kuliah
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

      {/*//! INFORMATIONAL */}
      <section className="position-relative mt-3">
        <Row className="row-gap-3">
          {/*//! materi pembelajaran*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Materi Pembelajaran
                </h3>
              </Col>
              <Col sm={12}>
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  text here
                </p>
              </Col>
            </Row>
          </Col>

          {/*//! Pustaka Utama*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Pustaka Utama
                </h3>
              </Col>
              <Col sm={12}>
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  text here
                </p>
              </Col>
            </Row>
          </Col>

          {/*//! Pustaka Pendukung*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Pustaka Pendukung
                </h3>
              </Col>
              <Col sm={12}>
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  text here
                </p>
              </Col>
            </Row>
          </Col>

          {/*//! Media Perangkat Lunak*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Media Perangkat Lunak
                </h3>
              </Col>
              <Col sm={12}>
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  text here
                </p>
              </Col>
            </Row>
          </Col>

          {/*//! Media Perangkat keras*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Media Perangkat keras
                </h3>
              </Col>
              <Col sm={12}>
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  text here
                </p>
              </Col>
            </Row>
          </Col>

          {/*//! Dokument RPS*/}
          <Col sm={12}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Dokument RPS
                </h3>
              </Col>
              <Col sm={12}>
                <div className="d-flex align-items-center gap-2">
                  {/* button file */}
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
                        File RPS.pdf
                      </Label>
                    </Button>
                  </div>

                  {/*  button delete */}
                  <Button color="transparent" className="text-danger p-0 ">
                    Hapus
                  </Button>
                </div>
              </Col>
            </Row>
          </Col>
        </Row>

        {/*//! BUTTON BACK  */}
        <section className="position-relative mt-3 d-flex justify-content-starts">
          <Link
            href={"?tab=cpl"}
            className="btn  d-flex align-items-center gap-2"
            color="transparent"
            style={{ color: "#10487A", border: "1px solid #10487A" }}
          >
            Kembali
          </Link>
        </section>
      </section>
    </section>
  );
}

export default SectionRPS;
