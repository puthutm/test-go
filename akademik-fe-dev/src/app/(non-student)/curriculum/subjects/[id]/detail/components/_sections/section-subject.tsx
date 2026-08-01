"use client";
import React from "react";
import { useParams } from "next/navigation";
import { signOut } from "next-auth/react";
// import third party component
import { Row, Col } from "reactstrap";
// import component
// import InformationSubject from "../information-subject";
import { DisabledByDefaultIcon } from "@/components/icons/disabled-by-default";
import { CheckBoxIcon } from "@/components/icons/checkbox";

import { useGetDetailSubjectLecturer } from "@/services/api/curriculum/lecture-subject/use-get-detail-subject-lecture";

function SectionSubject() {
  const params = useParams();

  const { data: dataSubjectDetail, isLoading: isLoadingDetailSubject } =
    useGetDetailSubjectLecturer(params.id as string);

  if (dataSubjectDetail?.status === 401) {
    signOut();
  }

  return (
    <section className="position-relative m-0 mt-2">
      <Row className="row-gap-3">
        {/*//! year curriculum */}
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
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.curriculum_year_name ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! unit pengampu */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Unit Pengampu
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.study_program_name ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! Kode Mata Kuliah */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Kode Mata Kuliah
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.code ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! Rumpun mata kuliah */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Rumpun Mata Kuliah
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {(dataSubjectDetail?.data.field_study_name as string) ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! nama mata kuliah id */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Nama Mata Kuliah {`(IND)`}
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.name_id ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! dosen pengampu */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Dosen Pengampu
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data?.supporting_lecturers == null
                    ? "-"
                    : dataSubjectDetail?.data?.supporting_lecturers.map(
                        (el) => {
                          return `${el.lecturer_front_title} ${el.lecturer_name} , ${el.lecturer_back_title}`;
                        }
                      )}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! nama mata kuliah eng */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Nama Mata Kuliah {`(EN)`}
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {(dataSubjectDetail?.data.name_en as string) ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! dosen pengembang rps */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Dosen Pengembang RPS
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data?.developer_rps_lecturers == null
                    ? "-"
                    : dataSubjectDetail?.data?.developer_rps_lecturers.map(
                        (el) => {
                          return `${el.lecturer_front_title} ${el.lecturer_name} , ${el.lecturer_back_title}`;
                        }
                      ) ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! JENIS mata kuliah */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Jenis Mata Kuliah
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {(dataSubjectDetail?.data.course_type_name as string) ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! Koord. pengampu mk */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Koord. Pengampu MK
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data?.subject_coordinator_lecturers ==
                  null
                    ? "-"
                    : dataSubjectDetail?.data?.subject_coordinator_lecturers.map(
                        (el) => {
                          return `${el.lecturer_front_title} ${el.lecturer_name} , ${el.lecturer_back_title}`;
                        }
                      ) ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! Kelompok Mata Kuliah */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Kelompok Mata Kuliah
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.course_group_name ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! MKU */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Apakah Termasuk MKU
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : dataSubjectDetail?.data?.is_mku ? (
                <CheckBoxIcon />
              ) : (
                <DisabledByDefaultIcon />
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SKS tatap muka */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                SKS Tatap Muka
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.face_to_face_sks ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SAP */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                SAP
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : dataSubjectDetail?.data?.is_sap ? (
                <CheckBoxIcon />
              ) : (
                <DisabledByDefaultIcon />
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SKS praktikum */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                SKS Praktikum
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.practicum_sks ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SILABUS */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Silabus
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : dataSubjectDetail?.data?.is_silabus ? (
                <CheckBoxIcon />
              ) : (
                <DisabledByDefaultIcon />
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SKS praktikum lapangan */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                SKS Praktikum Lapangan
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.field_practice_sks ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! bahan ajar */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Bahan Ajar
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : dataSubjectDetail?.data?.is_teaching_material ? (
                <CheckBoxIcon />
              ) : (
                <DisabledByDefaultIcon />
              )}
            </Col>
          </Row>
        </Col>

        {/*//! SKS Simulasi */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                SKS Simulasi
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.simulation_sks ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>

        {/*//! diktat */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Diktat
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : dataSubjectDetail?.data?.is_diktat ? (
                <CheckBoxIcon />
              ) : (
                <DisabledByDefaultIcon />
              )}
            </Col>
          </Row>
        </Col>
        {/*//! total sks */}
        <Col sm={6}>
          <Row className="row-gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#3A3A3A" }}
              >
                Total SKS
              </h3>
            </Col>
            <Col sm={12}>
              {isLoadingDetailSubject ? (
                <span className="placeholder w-100 " />
              ) : (
                <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                  {dataSubjectDetail?.data.total_sks ?? "-"}
                </p>
              )}
            </Col>
          </Row>
        </Col>
      </Row>
    </section>
  );
}

export default SectionSubject;
