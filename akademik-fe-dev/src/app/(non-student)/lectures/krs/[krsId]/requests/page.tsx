import { Card } from "reactstrap";

import { getDetailKRSRequest } from "@/services/api/lectures/krs/get-detail-krs-request";
import { StudentInfo } from "./components/student-info";
import { TableKrsDetail } from "./components/table-krs-detail";

export default async function KrsRequestPage({
  params,
}: {
  params: Promise<{ krsId: string }>;
}) {
  const krsId = (await params).krsId;

  const data = await getDetailKRSRequest(krsId);

  return (
    <Card>
      <div className="d-flex flex-column gap-3 p-3">
        <StudentInfo data={data} />
        <TableKrsDetail data={data} />
      </div>
    </Card>
  );
}
