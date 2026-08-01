import { notFound } from "next/navigation";
import { FormClassContract } from "./form-class-contract";

export default async function ClassContract({
  detailClass,
  params,
  isDetail,
}: {
  detailClass: ApiResponse<Class>;
  params: Promise<{ academicPeriodId: string; classId: string }>;
  isDetail?: boolean;
}) {
  const classId = (await params).classId;

  if (detailClass?.status === 404) return notFound();

  return (
    <div className="d-flex flex-column gap-2 pt-3 px-2">
      <FormClassContract
        classId={classId}
        detailClass={detailClass}
        isDetail={isDetail}
      />
    </div>
  );
}
