import { CloseIcon } from "@/components/icons/close";
import { EditIcon } from "@/components/icons/edit";
import styles from "@/styles/table-counseling.module.css";
import { EyeAkademikIcon } from "@/components/icons/eye-akademik";

import { useState } from "react";
import { Modal, Table } from "reactstrap";

import { InsertDriveFileIcon } from "@/components/icons/inser-drive-file";

export const CounselingTable = () => {
  const [modal, setModal] = useState(false);

  const toggle = () => setModal(!modal);
  return (
    <>
      <div className="table-responsive mt-3">
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
                style={{ maxWidth: "107px", backgroundColor: "#DEE5EC" }}
              >
                Waktu Bimbingan
              </th>
              <th scope="col" style={{ backgroundColor: "#DEE5EC" }}>
                Topik
              </th>
              <th scope="col" style={{ backgroundColor: "#DEE5EC" }}>
                Feedback Bimbingan
              </th>
              <th scope="col" style={{ backgroundColor: "#DEE5EC" }}>
                Status
              </th>
              <th scope="col" style={{ backgroundColor: "#DEE5EC" }}>
                Aksi
              </th>
            </tr>
          </thead>
          <tbody>
            <tr className="align-middle">
              <td className="text-center">2 Jan 2025 10:00 WIB</td>
              <td>Bab 1 Pendahuluan</td>
              <td>-</td>
              <td className="text-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor: "#DEE5EC",
                    color: "grey",
                    fontSize: "12px",
                  }}
                >
                  Menunggu Konfirmasi
                </span>
              </td>
              <td className="text-center">
                <div className="d-flex align-items-center gap-2">
                  <button
                    onClick={toggle}
                    style={{
                      border: "none",
                      padding: "0",
                      background: "none",
                    }}
                  >
                    <EyeAkademikIcon />
                  </button>
                  <CloseIcon height="20px" width="20px" />
                  <EditIcon height="20px" width="20px" />
                </div>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">27 Jan 2025 13:00 wib</td>
              <td>Bab 1 Pendahuluan</td>
              <td>Perlu menambahkan teori UCD User-...</td>
              <td className="text-center">
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
              <td className="text-center">
                <div className="d-flex align-items-center justify-content-center gap-2">
                  <button
                    onClick={toggle}
                    style={{
                      border: "none",
                      background: "none",
                    }}
                  >
                    <EyeAkademikIcon />
                  </button>
                </div>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">7 Feb 2025 14.00 WIB</td>
              <td>Bab 2 Kajian Pustaka</td>
              <td>Saya tidak bisa di waktu tsb. Tolong jadwalkan ulang</td>
              <td className="text-center">
                <span
                  className="px-2 py-1 rounded badge"
                  style={{
                    backgroundColor: "#C9C0BE",
                    color: "#A14835",
                    fontSize: "12px",
                  }}
                >
                  Ditolak
                </span>
              </td>
              <td className="text-center">
                <div className="d-flex align-items-center justify-content-center gap-2">
                  <button
                    onClick={toggle}
                    style={{
                      border: "none",
                      background: "none",
                    }}
                  >
                    <EyeAkademikIcon />
                  </button>
                </div>
              </td>
            </tr>
            <tr className="align-middle">
              <td className="text-center">7 Mar 2025 09:00 WIB</td>
              <td>Bab 2 Kajian Pustaka</td>
              <td>Perlu menambahkan teori UCD User-...</td>
              <td className="text-center">
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
              <td className="text-center">
                <div
                  className="d-flex align-items-center justify-content-center gap-2"
                  style={{
                    width: "100%",
                  }}
                >
                  <button
                    onClick={toggle}
                    style={{
                      border: "none",
                      background: "none",
                    }}
                  >
                    <EyeAkademikIcon />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </Table>
      </div>
      <Modal isOpen={modal} toggle={toggle} centered>
        <div className={`${styles.p_24}`}>
          <div
            className={`d-flex ${styles.h_25} gap-2 ${styles.mb_24} ${styles.w_100}`}
          >
            <div
              className={`d-flex justify-content-between align-items-center gap-2 ${styles.w_100}`}
            >
              <div className="d-flex align-items-center gap-2">
                <h2
                  className={`card-title fw-medium py-3 mb-0`}
                  style={{ color: "#495057" }}
                >
                  Bimbingan Kerja Praktik
                </h2>
                <span
                  className="rounded badge"
                  style={{
                    backgroundColor: "#6CBE401A",
                    color: "#6CBE40",
                    fontSize: "12px",
                  }}
                >
                  Disetujui
                </span>
              </div>
              <CloseIcon height="20px" width="20px" onClick={toggle} />
            </div>
          </div>
          <div
            className={`d-flex gap-2 align-items-center border-2 border-bottom `}
          ></div>
          <div className={`${styles.mt_24}`}>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>Waktu Bimbingan</p>
              <p className={`${styles.fSize14}`}>7 Feb 2025 14:00 WIB</p>
            </div>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>Dosen Pembimbing</p>
              <p className={`${styles.fSize14}`}>Dr. Rina Suryani, M.Kom.</p>
            </div>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>Tahap Saat ini</p>
              <p className={`${styles.fSize14}`}>Bab 2 Kajian Pustaka</p>
            </div>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>Deskripsi Bimbingan</p>
              <p className={`${styles.fSize14}`}>
                Pada bimbingan ini, saya ingin membahas dan meminta masukan
                terkait kerangka teori serta penelitian terdahulu yang relevan
                dengan pengembangan UI/UX dalam sistem akademik berbasis mobile.
                Saya telah menyusun literatur dari jurnal dan buku terkait
                usability, accessibility, dan user experience design, tetapi
                masih memerlukan arahan mengenai framework yang paling tepat
                untuk penelitian ini.
              </p>
            </div>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>File Dokumen</p>

              <button type="button" className={`${styles.btn_dokumen}`}>
                <InsertDriveFileIcon color="white" />
                <p>DraftTA-Insyaallah Fix</p>
              </button>
            </div>
            <div className={`${styles.mb_12}`}>
              <p className={`${styles.headingGrey}`}>Feedback Dosen</p>
              <p className={`${styles.fSize14}`}>
                Perlu menambahkan teori UCD (User-Centered Design) sebagai dasar
                pendekatan pengembangan UI/UX. Cek juga penelitian tentang
                gamifikasi dalam sistem akademik untuk meningkatkan engagement
                pengguna.
              </p>
            </div>
          </div>
        </div>
      </Modal>
    </>
  );
};
