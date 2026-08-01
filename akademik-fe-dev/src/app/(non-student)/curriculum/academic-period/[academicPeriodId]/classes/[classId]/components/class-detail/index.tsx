import { FormClassDetail } from "./form-class-detail";

export default async function ClassDetail({
  detailClass,
  params,
  isDetail,
}: {
  detailClass: ApiResponse<Class>;
  params: Promise<{ academicPeriodId: string; classId: string }>;
  isDetail?: boolean;
}) {
  const classId = (await params).classId;

  return (
    <FormClassDetail data={detailClass} classId={classId} isDetail={isDetail} />
  );
}
