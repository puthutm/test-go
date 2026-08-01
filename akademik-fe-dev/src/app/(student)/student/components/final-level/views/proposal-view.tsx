"use client";

import { useGetFinalProjectProposal } from "@/services/api/students/final-project-proposal/use-get-final-project-project-proposal";
import { EmptyStateProposal } from "../empty-state-proposal";
import { ModalUploadProposal } from "../modal-upload-proposal";
import { TableProposal } from "../table-proposal";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { useModalContext } from "@/lib/hooks/use-modal";
import { Button } from "reactstrap";

export default function ProposalFinalLevelView() {
  const { setModalState } = useModalContext();
  const { data, isLoading, refetch } = useGetFinalProjectProposal();

  if (isLoading) {
    return (
      <div className="d-flex justify-content-center align-items-center">
        <div className="spinner-border text-primary" role="status">
          <span className="visually-hidden">Loading...</span>
        </div>
      </div>
    );
  }

  if (data?.error) {
    return (
      <div className="d-flex flex-column gap-2">
        <p className="fs-4 pb-0">Something went wrong!</p>
        <Button
          color="primary"
          size="md"
          onClick={() => refetch()}
          style={{ width: "fit-content" }}
        >
          Try again
        </Button>
      </div>
    );
  }

  return (
    <>
      <div className="d-flex justify-content-between align-items-center border-2 border-bottom  py-3">
        <h2
          className={`card-title fw-medium fs-5 mb-0 `}
          style={{ color: "#495057" }}
        >
          Proposal Tugas Akhir
        </h2>
        {data?.data.length ? (
          <button
            className="btn d-flex align-items-center gap-2 p-2 text-primary"
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
            <FileDownloadIcon /> Upload Proposal
          </button>
        ) : (
          ""
        )}
      </div>
      <ModalUploadProposal />
      {data?.data.length ? (
        <TableProposal data={data?.data as FinalProjectProposal[]} />
      ) : (
        <EmptyStateProposal />
      )}
    </>
  );
}
