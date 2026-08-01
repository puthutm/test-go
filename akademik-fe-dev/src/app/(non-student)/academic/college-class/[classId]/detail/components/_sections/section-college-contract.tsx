"use client";
import React from "react";
import { useParams } from "next/navigation";
// import third party component
import { Card, CardBody, CardHeader, Row, Col } from "reactstrap";

// import component
import TabsSectionDetailCollegeClass from "../tab-detail-college-class";
import InformationCollegeClass from "../information-college-class";

import { useGetCourseContract } from "@/services/api/academic/lecturer/class-schedule/detail-class/course-contract/use-get-course-contract";

import { formatDate } from "@/lib/utils/format-date";


function SectionCollegeContract() {
  const params = useParams();
  const { data: dataCourseContract, isLoading: isLoadingCourseContract } =
    useGetCourseContract(params.classId as string);

  return (
    <section className="position-relative">
      {/*//! INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0  border-0">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 mt-1">
              Kontrak Kuliah
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

      {/*//! CONTENT */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        <Row className="row-gap-3">
          {/*//! periode akademik */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Periode Akademik
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.academic_periode_fullname ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Kapasitas */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Kapasitas
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.capacity ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Program Studi */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Program Studi
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.study_program_name ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Tanggal Mulai */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Tanggal Mulai
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {formatDate(
                      dataCourseContract?.data.start_date_of_college ?? "-"
                    )}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Tahun Kurikulum */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Tahun Kurikulum
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.curriculum_year_name ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Tanggal Selesai */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Tanggal Selesai
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {formatDate(
                      dataCourseContract?.data.end_date_of_college ?? "-"
                    )}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Mata Kuliah */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Mata Kuliah
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.subject_name_en ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Jumlah Pertemuan */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Jumlah Pertemuan
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.number_of_meeting ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>

          {/*//! Nama Kelas */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Nama Kelas
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.name ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>
          {/*//! Sistem Kuliah */}
          <Col sm={6}>
            <Row className="row-gap-1">
              <Col sm={12}>
                <h3
                  className="m-0 p-0 fw-semibold fs-6"
                  style={{ color: "#3A3A3A" }}
                >
                  Sistem Kuliah
                </h3>
              </Col>
              <Col sm={12}>
                {isLoadingCourseContract ? (
                  <span className="placeholder w-100 " />
                ) : (
                  <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                    {dataCourseContract?.data.lecturer_system ?? "-"}
                  </p>
                )}
              </Col>
            </Row>
          </Col>
        </Row>
      </Card>
    </section>
  );
}

export default SectionCollegeContract;
//   {/*//! Dokument */}
//   <Col sm={12}>
//     <Row className="row-gap-1">
//       <Col sm={12}>
//         <h3
//           className="m-0 p-0 fw-semibold fs-6"
//           style={{ color: "#3A3A3A" }}
//         >
//           Dokumen Kontrak Kuliah
//         </h3>
//       </Col>
//       <Col sm={12}>
//         <div className="d-flex align-items-center gap-2">
//           {/* button file */}
//           <div className="">
//             <input
//               // ref={inputImport}
//               type="file"
//               // onChange={handleImport}
//               style={{ display: "none" }}
//               name="fImport"
//               id="Fimport"
//               accept=".xlsx, .xls, .csv"
//             />
//             <Button
//               type="button"
//               // disabled={isFetchingExportUser || isPendingImportUser}
//               className="btn create-btn d-flex align-items-center p-0 text-light"
//               style={{ background: "#44A7FF" }}
//               // onClick={handleShow}
//               color="transparent"
//               data-bs-target="#api-key-modal"
//             >
//               <Label
//                 htmlFor="Fimport"
//                 className="py-1 px-3 m-0 pointer fs-6"
//                 style={{ cursor: "pointer" }}
//               >
//                 DokumenKontrakKuliah.pdf
//               </Label>
//             </Button>
//           </div>
//         </div>
//       </Col>
//     </Row>
//   </Col>
