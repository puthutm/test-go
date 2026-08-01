import React, { Suspense } from "react";

import { getAllClassScore } from "@/services/api/settings/academic-period/class-score/get-all-class-score";
import { TableClassScore } from "./table-class-score";
import { getAllGradeCompositions } from "@/services/api/settings/grade-composition/get-all-grade-composition";
import { Spinner } from "reactstrap";
import { checkOpenCloseGradeByClass } from "@/services/api/settings/academic-period/open-close-grade/check-open-close-grade-by-class";

export default async function ClassScoreView({
  searchParams,
  detailClass,
}: {
  searchParams: { [key: string]: string | string[] | undefined };
  detailClass: ApiResponse<Class>;
}) {
  const academicPeriodId = detailClass?.data.academic_periode_id;
  const classId = detailClass?.data.id;

  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";

  const [data, gradeComposition, openCloseClassScore] = await Promise.all([
    getAllClassScore({
      academicPeriodId,
      classId,
      queryParam: {
        page,
        search: search as string,
      },
    }),
    getAllGradeCompositions({
      page: 1,
      limit: 100,
    }),
    checkOpenCloseGradeByClass({ academicPeriodId, classId }),
  ]);

  return (
    <Suspense
      fallback={
        <div>
          <Spinner />
        </div>
      }
    >
      <TableClassScore
        data={data}
        gradeComposition={gradeComposition}
        statusLockClassScore={Boolean(openCloseClassScore?.data?.status_lock)}
      />
    </Suspense>
  );
}
