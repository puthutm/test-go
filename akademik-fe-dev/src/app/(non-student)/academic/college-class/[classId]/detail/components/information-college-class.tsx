"use client";
import React from "react";
import { useParams } from "next/navigation";
import { useSearchParams } from "next/navigation";
import { signOut } from "next-auth/react";
import { Row, Col } from "reactstrap";
import { useGetDetailClassScheduleLecturer } from "@/services/api/academic/lecturer/class-schedule/detail-class/use-detail-schedule-academic-lecturer";
import { formatDate } from "@/lib/utils/format-date";

function InformationCollegeClass() {
  const params = useParams();
  const searchQuery = useSearchParams();

  const {
    data: dataDetailClassSchedule,
    isFetching: isLoadingDetailClassSchedule,
  } = useGetDetailClassScheduleLecturer(params.classId as string);

  if (dataDetailClassSchedule?.status === 401) {
    signOut();
  }
  return (
    <div className="p-3 rounded-3" style={{ background: "#FAFCFF" }}>
      <Row className="row-gap-3">
        {
          //? information for detail class
          searchQuery.get("tab") === "class" ? (
            <>
              {/*//! periode akademik  */}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data
                          .academic_periode_fullname ?? "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.capacity ?? "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.study_program_name ??
                          "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {formatDate(
                          dataDetailClassSchedule?.data.start_date_of_college ??
                            "-"
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.curriculum_year_name ??
                          "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {formatDate(
                          dataDetailClassSchedule?.data.end_date_of_college ??
                            "-"
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.subject_name_en ?? "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.number_of_meeting ?? "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.name ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
              {/*//!MBKM */}
              {/* <Col sm={6}>
                <Row className="row-gap-1">
                  <Col sm={12}>
                    <h3
                      className="m-0 p-0 fw-semibold fs-6"
                      style={{ color: "#3A3A3A" }}
                    >
                      MBKM Kampus Merdeka
                    </h3>
                  </Col>
                  <Col sm={12}>
                    <CheckBoxIcon/> 
                    <DisabledByDefaultIcon />
                  </Col>
                </Row>
              </Col> */}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.lecturer_system ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
            </>
          ) : (
            //? information for detail class else
            <>
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.study_program_name ??
                          "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
              {/*//! periode akademik  */}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data
                          .academic_periode_fullname ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
              {/*//! name class  */}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data?.name ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
              {/*//! kurikulum  */}
              <Col sm={6}>
                <Row className="row-gap-1">
                  <Col sm={12}>
                    <h3
                      className="m-0 p-0 fw-semibold fs-6"
                      style={{ color: "#3A3A3A" }}
                    >
                      Kurikulum
                    </h3>
                  </Col>
                  <Col sm={12}>
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data?.curriculum_year_name ??
                          "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.lecturer_system ?? "-"}
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
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.capacity ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
              {/*//! peserta */}
              <Col sm={6}>
                <Row className="row-gap-1">
                  <Col sm={12}>
                    <h3
                      className="m-0 p-0 fw-semibold fs-6"
                      style={{ color: "#3A3A3A" }}
                    >
                      Peserta
                    </h3>
                  </Col>
                  <Col sm={12}>
                    {isLoadingDetailClassSchedule ? (
                      <span className="placeholder w-100 " />
                    ) : (
                      <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                        {dataDetailClassSchedule?.data.total_participant ?? "-"}
                      </p>
                    )}
                  </Col>
                </Row>
              </Col>
            </>
          )
        }
      </Row>
    </div>
  );
}

export default InformationCollegeClass;
