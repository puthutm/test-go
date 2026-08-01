import Link from "next/link";
import React from "react";
import { Table } from "reactstrap";

const RequirementTabContent = () => {
  const data = {
    title: "Pengembangan UI/UX Sistem Akademik Berbasis Mobile",
    requirements: [
      {
        title: "Biaya Ujian TA — Bukti Pelunasan UKT",
        info: "Bukti pelunasan UKT",
        file: "http://google.co.id",
        status: "pending",
      },
      {
        title:
          "Jumlah Bimbingan — Jumlah bimbingan mahasiswa minimal 5 kali per dosen pembimbing",
        info: "Jumlah bimbingan mahasiswa minimal 5 kali per dosen pembimbing",
        file: "http://google.co.id",
        status: "approved",
      },
      {
        title: "Laporan — Laporan Tugas Akhir",
        info: "Laporan Tugas Akhir",
        file: "http://google.co.id",
        status: "rejected",
      },
      {
        title: "Lulus SKS — Lulus 144 SKS (Dibuktikan dengan KRS & KHS)",
        info: "Lulus 144 SKS (Dibuktikan dengan KRS & KHS)",
        file: "http://google.co.id",
        status: "approved",
      },
      {
        title:
          "Persetujuan Pembimbing — Form Persetujuan Sidang dari Dosen Pembimbing",
        info: "Form Persetujuan Sidang dari Dosen Pembimbing",
        file: "http://google.co.id",
        status: "rejected",
      },
    ],
  };

  return (
    <div>
        <div className="border-bottom border-3 mb-2">
        <h5 className="fw-semibold">Syarat Ujian</h5>
      </div>
      <div>
        <p>Judul Tugas Akhir:</p>
        <p className="fw-semibold text-primary">{data.title}</p>
      </div>

      <div className="table-responsive mt-3">
        <Table striped className="table-bordered">
          <thead className="table-light">
            <tr>
              <th className="text-center">No</th>
              <th className="text-center">Syarat Ujian</th>
              <th className="text-center">Keterangan</th>
              <th className="text-center">Berkas</th>
              <th className="text-center">Status</th>
              {/* <th className="text-center">Aksi</th> */}
            </tr>
          </thead>

          <tbody>
            {data.requirements.map((item, index) => (
              <tr key={index}>
                <td className="text-center">{index + 1}</td>
                <td>{item.title}</td>
                <td>{item.info}</td>
                <td>
                  <Link
                    href={item.file}
                    target="_blank"
                    className="btn btn-primary"
                  >
                    Berkas
                  </Link>
                </td>
                <td className="text-center">
                  {item.status === "pending" && (
                    <span className="badge bg-warning">Pending</span>
                  )}
                  {item.status === "approved" && (
                    <span className="badge bg-success">Approved</span>
                  )}
                  {item.status === "rejected" && (
                    <span className="badge bg-danger">Rejected</span>
                  )}
                </td>
                {/* <td className="text-center">
                  <i className="mdi mdi-eye-outline"></i>
                </td> */}
              </tr>
            ))}
          </tbody>
        </Table>
      </div>
    </div>
  );
};

export default RequirementTabContent;
