"use client";
import { useState } from "react";
import { CounselingTable } from "../table-counseling";
import { EmptyStateCounselingKp } from "../empty-state-counseling-kp";

export default function CounselingKpViews() {
  const [hasBimbangan] = useState(true);
  return (
    <div>
      <div className="d-flex justify-content-between align-items-center border-2 border-bottom">
        <h2
          className={`card-title fw-medium py-3 mb-0`}
          style={{ color: "#495057" }}
        >
          Bimbingan Kerja Praktik
        </h2>
      </div>

      {hasBimbangan ? <CounselingTable /> : <EmptyStateCounselingKp />}
    </div>
  );
}
