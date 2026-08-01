import Link from "next/link";

import { HighlightOffIcon } from "@/components/icons/highlight-off";
import { OpenInNewIcon } from "@/components/icons/open-in-new";

export const EmptyStateCounselingFinalTask = () => {
  return (
    <div className="py-4 d-flex justify-content-center align-items-center gap-3 flex-column px-2">
      <HighlightOffIcon />
      <div className="d-flex justify-content-center align-items-center gap-2 flex-column text-muted">
        <h2 className="fw-semibold fs-5 mb-0">
          Belum Ada Informasi Bimbingan Tugas Akhir
        </h2>
        <p className="font-normal fs-6">
          Anda belum membuat bimbingan tugas akhir. Silakan atur jadwal
          bimbingan dengan dosen pembimbing melalui LMS.
        </p>
      </div>
      <Link
        className="w-100 btn btn-primary"
        href="https://lms.unsia.ac.id"
        target="_blank"
      >
        Pergi ke LMS <OpenInNewIcon />
      </Link>
    </div>
  );
};
