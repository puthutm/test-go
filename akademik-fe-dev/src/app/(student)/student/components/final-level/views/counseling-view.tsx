import { EmptyStateCounselingFinalTask } from "../empty-state-counseling-final-level";

export default function CounselingView() {
  const hasCounseling = true;
  return (
    <>
      <div className="d-flex align-items-center py-3 mb-0 border-2 border-bottom gap-2">
        <h2
          className={`card-title fw-medium fs-5 mb-0`}
          style={{ color: "#495057" }}
        >
          Bimbingan Tugas Akhir
        </h2>
      </div>
      {!hasCounseling ? <></> : <EmptyStateCounselingFinalTask />}
    </>
  );
}
