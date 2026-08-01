import { ErrorIcon } from "@/components/icons/error";

export default function JudiciaryView() {
  return (
    <>
      <div className="d-flex align-items-center py-3 mb-0 border-2 border-bottom gap-2">
        <h2
          className={`card-title fw-medium fs-5 mb-0`}
          style={{ color: "#495057" }}
        >
          Yudisium
        </h2>
      </div>
      <div className="mt-3">
        <div
          className="alert alert-danger alert-border-left alert-dismissible fade show py-1"
          role="alert"
        >
          <ErrorIcon color="#921A00" />
          <span style={{ color: "#921A00" }}>
            Anda belum terdaftar menjadi Peserta Yusidium, Harap konfirmasi ke
            bagian Adminstrasi
          </span>
        </div>
      </div>
    </>
  );
}
