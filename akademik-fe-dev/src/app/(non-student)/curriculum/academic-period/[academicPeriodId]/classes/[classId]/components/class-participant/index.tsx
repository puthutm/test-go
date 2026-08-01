import { notFound } from "next/navigation";

import { TableClassParticipant } from "./table-class-participant";
import { getClassParticipantForProgramHead } from "@/services/api/curriculum/academic-period/class-participant/get-class-participant";

export default async function ClassParticipant({
  detailClass,
  searchParams,
  isDetail,
}: {
  detailClass: ApiResponse<Class>;
  searchParams: { [key: string]: string | string[] | undefined };
  isDetail?: boolean;
}) {
  const studyProgramId = detailClass?.data.study_program_id;
  const classId = detailClass?.data.id;

  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const data = await getClassParticipantForProgramHead(classId, {
    page,
    search: search as string,
  });

  if (data.status === 401) {
    return notFound();
  }

  return (
    <div className="d-flex flex-column gap-3">
      <p className="fw-medium fs-5" style={{ color: "#3A3A3A" }}>
        Peserta Kelas
      </p>
      <TableClassParticipant
        studyProgramId={studyProgramId}
        classId={classId}
        participant={data}
        isDetail={isDetail}
      />
    </div>
  );
}
