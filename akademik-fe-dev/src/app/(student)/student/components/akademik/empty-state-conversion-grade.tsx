import { HighlightOffIcon } from "@/components/icons/highlight-off";

export const EmptyStateConversionGrade = () => {
  return (
    <div className="py-4 d-flex justify-content-center align-items-center gap-3 flex-column px-2">
      <HighlightOffIcon />
      <div className="d-flex justify-content-center align-items-center gap-2 flex-column text-muted">
        <h2 className="fw-semibold fs-5 mb-0">
          Belum Ada Informasi Nilai Konversi
        </h2>
        <p className="font-normal fs-6">
          Silakan cek kembali nanti atau hubungi pihak akademik untuk informasi
          lebih lanjut.
        </p>
      </div>
    </div>
  );
};
