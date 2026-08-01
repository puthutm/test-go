import { ErrorIcon } from "@/components/icons/error";

export default function GraduationView() {
  return (
    <>
      <div className="d-flex align-items-center py-3 mb-0 border-2 border-bottom gap-2">
        <h2
          className={`card-title fw-medium fs-5 mb-0`}
          style={{ color: "#495057" }}
        >
          Wisuda
        </h2>
      </div>
      <div className="mt-3">
        <div
          className="alert alert-danger alert-border-left alert-dismissible fade show py-1"
          role="alert"
        >
          <ErrorIcon color="#921A00" />{" "}
          <span className="fw-semibold" style={{ color: "#921A00" }}>
            18/430862/SA/19477 - Haerunnisa - Belum Terjadwal
          </span>
        </div>
      </div>
    </>
  );
}
