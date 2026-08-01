"use client";

import { ErrorIcon } from "@/components/icons/error";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { ForwardToInboxIcon } from "@/components/icons/forward-to-inbox";
import { semesters } from "@/lib/constants/table-khs-data";
import { useGetKhs } from "@/services/api/students/academic/khs/use-get-khs";
import React, { useState } from "react";
import { Spinner } from "reactstrap";

const TableKHS = () => {
  const [semesterFilter, setSemesterFilter] = useState<string>("Semua");
  const [showEmailAlert, setShowEmailAlert] = useState<boolean>(false);
  const [showDownloadAlert, setShowDownloadAlert] = useState<boolean>(false);

  const { data, isLoading } = useGetKhs();

  if (isLoading) {
    return (
      <div className="d-flex justify-content-center py-5">
        <Spinner />
      </div>
    );
  }

  // Calculate total SKS from filtered semesters (or all if no filter)
  const totalSKSCumulative = data?.data?.semesters.reduce(
    (sum, semester) => sum + semester.total_sks,
    0
  );

  const handleSendEmail = () => {
    setShowEmailAlert(true);
    setTimeout(() => setShowEmailAlert(false), 3000);
  };

  const handleDownloadKHS = () => {
    setShowDownloadAlert(true);
    setTimeout(() => setShowDownloadAlert(false), 3000);
  };

  return (
    <>
      <div className="d-flex justify-content-between align-items-center mb-4 border-bottom border-2 ">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Hasil Studi
        </h5>
        <div className="d-flex gap-2">
          <select
            className="form-select d-flex align-items-center"
            value={semesterFilter}
            onChange={(e) => setSemesterFilter(e.target.value)}
          >
            <option value="Semua">Semua</option>
            {semesters.map((semester) => (
              <option
                key={`${semester.semester}`}
                value={`${semester.semester}`}
              >
                Semester {semester.semester} - {semester.year}/{semester.period}
              </option>
            ))}
          </select>
          <button
            onClick={handleSendEmail}
            className="btn d-flex align-items-center gap-2 text-primary"
            style={{
              whiteSpace: "nowrap",
              border: "1px solid #10487A",
              backgroundColor: "transparent",
            }}
          >
            <ForwardToInboxIcon /> Kirim ke Email
          </button>
          <button
            onClick={handleDownloadKHS}
            className="btn d-flex align-items-center gap-2 text-primary"
            style={{
              whiteSpace: "nowrap",
              border: "1px solid #10487A",
              backgroundColor: "transparent",
            }}
          >
            <FileDownloadIcon /> Download KHS
          </button>
        </div>
      </div>

      {/* Success alert for email */}
      {showEmailAlert && (
        <div className="alert alert-success" role="alert">
          <strong>Yey! Everything worked! </strong> Email berhasil dikirim!
          —check it out!
        </div>
      )}

      {/* Success alert for download */}
      {showDownloadAlert && (
        <div className="alert alert-success" role="alert">
          <strong>Yey! Everything worked! </strong> KHS berhasil diunduh! —check
          it out!
        </div>
      )}

      {/* Warning message */}
      {/* <div className="alert alert-warning alert-border-left" role="alert">
        Silakan hubungi Operator Sistem Informasi Akademik jika terdapat
        perbedaan data mata kuliah, SKS, atau nilai.
      </div> */}

      {/* Error message */}
      <div
        className="alert alert-danger alert-border-left d-inline-flex gap-2"
        role="alert"
      >
        <ErrorIcon color="#921A00" />
        <div>
          <p className="mb-0">
            Terdapat mata kuliah dengan nilai di bawah rata-rata. Anda perlu
            mengulang untuk memenuhi syarat kelulusan
          </p>
          <ul className="mt-3">
            {data?.data?.not_passed.map((item, index) => (
              <li className="mb-0" key={item.academic_periode_id + index}>
                {`${item.subject_code} - ${item.subject_name} / ${item.academic_periode_name}`}
              </li>
            ))}
          </ul>
        </div>
      </div>

      {/* Tables for each semester */}
      {data?.data?.semesters.map((semester, index) => (
        <div key={semester.academic_periode_name + index} className="mb-4">
          <h2
            className="h5 fw-semibold py-3 px-2 m-0"
            style={{ backgroundColor: "#FFE91D33" }}
          >
            {semester.academic_periode_name}
          </h2>
          <div className="table-responsive">
            <table className="table table-nowrap">
              <thead className="table-light">
                <tr>
                  <th scope="col">No</th>
                  <th scope="col">Kode</th>
                  <th scope="col">Mata Kuliah</th>
                  <th scope="col">Dosen</th>
                  <th scope="col">SKS</th>
                  <th scope="col">Nilai</th>
                  <th scope="col">Bobot</th>
                </tr>
              </thead>
              <tbody>
                {semester.subjects.map((subject, index) => (
                  <tr key={subject.subject_code + index}>
                    <td>{index + 1}</td>
                    <td>{subject.subject_code}</td>
                    <td>{subject.subject_name}</td>
                    <td>{"-"}</td>
                    <td>{subject.total_sks}</td>
                    <td>{subject.grade_code}</td>
                    <td>{subject.weight}</td>
                  </tr>
                ))}
                <tr className="table-light">
                  <td className="fw-medium mb-0 border-0">Total SKS</td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                  <td className="fw-medium mb-0 border-0">
                    {semester.total_sks}
                  </td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                </tr>
                <tr className="table-light">
                  <td className="fw-medium mb-0 border-0">IPS</td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                  <td className="fw-medium mb-0 border-0">
                    {semester.total_sks}
                  </td>
                  <td className="border-0"></td>
                  <td className="border-0"></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      ))}

      {/* Cumulative Total SKS */}
      <div className="bg-primary text-white p-3">
        <p className="fw-medium mb-0">
          Total SKS Kumulatif: {totalSKSCumulative}
        </p>
      </div>
    </>
  );
};

export default TableKHS;
