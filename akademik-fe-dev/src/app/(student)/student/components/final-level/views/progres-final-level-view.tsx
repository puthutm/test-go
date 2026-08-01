import { EmptyStateProgressFinalTask } from "../empty-state-progress-final-level";

export default function ProgressFinalTaskView() {
  const hasFinalTask = true;
  return (
    <>
      <div className="d-flex align-items-center py-3 mb-0 border-2 border-bottom gap-2">
        <h2
          className={`card-title fw-medium fs-5 mb-0`}
          style={{ color: "#495057" }}
        >
          Tugas Akhir
        </h2>
      </div>
      {!hasFinalTask ? <></> : <EmptyStateProgressFinalTask />}
    </>
  );
}
