"use client";

import { FileDownloadIcon } from "@/components/icons/file-download";
import { Col, Row, Table } from "reactstrap";

export const TableKrs = () => {
  const semesters = ["Semua", "Semester 1", "Semester 2", "Semester 3"];
  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Kartu Rencana Studi
        </h5>
        <div className="d-flex gap-2">
          <select className="form-select py-1 px-3">
            {semesters.map((semester) => (
              <option key={`${semester}`} value={`${semester}`}>
                {semester}
              </option>
            ))}
          </select>
          <button
            className="btn d-flex align-items-center gap-2 text-primary"
            style={{
              whiteSpace: "nowrap",
              border: "1px solid #10487A",
              backgroundColor: "transparent",
            }}
            onClick={() => alert("Downloaded")}
          >
            <FileDownloadIcon /> Download KRS
          </button>
        </div>
      </div>
      <h2
        className="py-3"
        style={{ color: "#909090", fontSize: "14px", fontWeight: "600" }}
      >
        Semester 5
      </h2>
      <div className="table-responsive">
        <Table
          borderless
          hover
          style={{ tableLayout: "auto" }}
          className="align-center"
        >
          <thead
            className="table-light text-center"
            style={{ backgroundColor: "#DEE5EC" }}
          >
            <tr className="align-middle">
              <th
                scope="col"
                style={{
                  maxWidth: "126px",
                  backgroundColor: "#DEE5EC",
                  color: "#495057",
                }}
              >
                No
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Kode MK
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Mata Kuliah
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                SKS
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Kelas
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Dosen
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Status
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="align-middle">
              <td className="text-center">1</td>
              <td>PB0104</td>
              <td>Algoritma & Pemrograman II</td>
              <td>2</td>
              <td>301</td>
              <td>Vika febri muliati, S.KOM, M.Kom</td>
              <td className="text-center align-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor: "#F065481A",
                    color: "#F06548",
                    fontSize: "12px",
                  }}
                >
                  Ditolak
                </span>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">2</td>
              <td>PB0104</td>
              <td>Algoritma & Pemrograman II</td>
              <td>2</td>
              <td>301</td>
              <td>Vika febri muliati, S.KOM, M.Kom</td>
              <td className="text-center align-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor: "#6CBE401A",
                    color: "#6CBE40",
                    fontSize: "12px",
                  }}
                >
                  Disetujui
                </span>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">3</td>
              <td>PB0104</td>
              <td>Algoritma & Pemrograman II</td>
              <td>2</td>
              <td>301</td>
              <td>Vika febri muliati, S.KOM, M.Kom</td>
              <td className="text-center align-center">
                <span
                  className="badge px-2 py-1 rounded w-100"
                  style={{
                    backgroundColor: "#F7B84B3B",
                    color: "#F7B84B",
                    fontSize: "12px",
                  }}
                >
                  Menunggu Persetujuan
                </span>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">4</td>
              <td>PB0104</td>
              <td>Algoritma & Pemrograman II</td>
              <td>2</td>
              <td>301</td>
              <td>Vika febri muliati, S.KOM, M.Kom</td>
              <td className="text-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor: "#C9C0BE",
                    color: "#A14835",
                    fontSize: "12px",
                  }}
                >
                  Revisi
                </span>
              </td>
            </tr>
          </tbody>
        </Table>
      </div>
      <div className="py-3 px-3" style={{ backgroundColor: "#DEE5EC" }}>
        <Row>
          <Col sm={11}>Total SKS yang Diambil</Col>
          <Col sm={1} className="pe-0">
            20/20
          </Col>
        </Row>
        <Row className="mt-2 fw-semibold">
          <Col sm={11}>Total SKS yang Disetujui</Col>
          <Col sm={1} className="pe-0">
            18/20
          </Col>
        </Row>
      </div>
    </>
  );
};
