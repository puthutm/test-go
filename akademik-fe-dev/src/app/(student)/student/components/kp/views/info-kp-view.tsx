"use client";

import { EmptyStateInfoKp } from "../empty-state-info-kp";
import { useState } from "react";
import { EditIcon } from "@/components/icons/edit";
import { AddIcon } from "@/components/icons/add";
import { FormInfoKp } from "../form-info-kp";

export default function InfoKpView() {
  const [isEdit, setIsEdit] = useState(false);
  const hasProposal = false;
  const hasPembimbing = true;

  return (
    <>
      <div className="d-flex justify-content-between align-items-center border-2 border-bottom">
        <h2
          className={`card-title fw-medium py-3 mb-0`}
          style={{ color: "#495057" }}
        >
          Proposal Kerja Praktik
        </h2>
        {hasProposal && hasPembimbing ? (
          <button
            className="bg-transparent rounded px-3 d-flex gap-1 align-items-center justify-content-center text-primary"
            style={{ border: "1px solid #10487A", paddingBlock: "8px" }}
            onClick={() => setIsEdit(true)}
          >
            <EditIcon />
            <span>Edit</span>
          </button>
        ) : null}
      </div>
      {hasPembimbing ? (
        <div className="d-flex flex-column mt-3">
          <span
            className="fw-semibold"
            style={{
              color: "#909090",
              fontSize: "12px",
              letterSpacing: "0.06px",
            }}
          >
            Dosen Pembimbing
          </span>
          <p
            style={{
              color: "#495057",
              fontSize: "14px",
              fontWeight: "500",
            }}
          >
            Vika febri muliati, S.KOM, M.Kom
          </p>
          {isEdit ? (
            <FormInfoKp
              hasProposal={hasProposal}
              isEdit={isEdit}
              setIsEdit={setIsEdit}
            />
          ) : (
            <button
              className="w-100 btn btn-primary mt-4"
              onClick={() => setIsEdit(true)}
            >
              <AddIcon color="white" /> Tambah Informasi Kerja Praktik
            </button>
          )}
        </div>
      ) : (
        <EmptyStateInfoKp />
      )}
    </>
  );
}
