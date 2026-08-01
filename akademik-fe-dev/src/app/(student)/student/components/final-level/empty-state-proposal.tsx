"use client";

import { FileUploadIcon } from "@/components/icons/file-upload";
import { HighlightOffIcon } from "@/components/icons/highlight-off";
import { useModalContext } from "@/lib/hooks/use-modal";

export const EmptyStateProposal = () => {
  const { setModalState } = useModalContext();
  return (
    <div className="py-4 d-flex justify-content-center align-items-center gap-3 flex-column px-2">
      <HighlightOffIcon />
      <div className="d-flex justify-content-center align-items-center gap-2 flex-column text-muted">
        <h2 className="fw-semibold fs-5 mb-0">
          Belum Ada Informasi Proposal Tugas Akhir
        </h2>
        <p className="font-normal fs-6">
          Anda belum mengupload proposal tugas akhir. Silakan upload proposal
          pertama Anda.
        </p>
      </div>
      <button
        className="w-100 btn btn-primary"
        onClick={() =>
          setModalState((prev) => ({
            ...prev,
            open: true,
          }))
        }
      >
        <FileUploadIcon color="white" /> Upload Proposal
      </button>
    </div>
  );
};
