"use client";

import { AddIcon } from "@/components/icons/add";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalContext } from "@/lib/hooks/use-modal";
import { Table } from "reactstrap";
import { ModalLetter } from "./modal-letter";

export const TableLetter = () => {
  const { setModalState } = useModalContext();
  return (
    <>
      <ModalLetter />
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Kartu Rencana Studi
        </h5>
        <button
          className="btn d-flex align-items-center gap-2 text-primary"
          style={{
            whiteSpace: "nowrap",
            border: "1px solid #10487A",
            backgroundColor: "transparent",
          }}
          onClick={() =>
            setModalState((prev) => ({
              ...prev,
              open: true,
              state: "add",
            }))
          }
        >
          <AddIcon color="#10487A" /> Ajukan Surat
        </button>
      </div>

      <div className="table-responsive mt-4">
        <Table
          borderless
          hover
          style={{ tableLayout: "auto" }}
          className="align-center"
        >
          <thead className="table-light text-center">
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
                Jenis Surat
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Tanggal Pengajuan
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Status
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Keterangan
              </th>
              <th
                scope="col"
                style={{ backgroundColor: "#DEE5EC", color: "#495057" }}
              >
                Aksi
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="align-middle">
              <td className="text-center">1</td>
              <td>PB0104</td>
              <td>Algoritma & Pemrograman II</td>
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
              <td>
                Lorem ipsum dolor, sit amet consectetur adipisicing elit. A,
                provident!
              </td>
              <td>
                <div className="d-flex gap-1 align-items-center">
                  <VisibilityIcon />
                  <FileDownloadIcon />
                </div>
              </td>
            </tr>
          </tbody>
        </Table>
      </div>
    </>
  );
};
