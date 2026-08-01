import { HighlightOffIcon } from "@/components/icons/highlight-off";

export const EmptyStateInfoKp = () => {
  return (
    <div className="py-4 d-flex justify-content-center align-items-center gap-3 flex-column px-2">
      <HighlightOffIcon />
      <div className="d-flex justify-content-center align-items-center gap-2 flex-column text-muted">
        <h2 className="fw-semibold fs-5 mb-0">
          Belum Ada Informasi Proposal Kerja Praktik
        </h2>
        <p className="font-normal fs-6">
          Anda belum mendapatkan dosen pembimbing. Mohon cek secara berkala atau
          hubungi Koordinator Kerja Praktik.
        </p>
      </div>
    </div>
  );
};
