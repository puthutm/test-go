import { getClassScheduleTemplate } from "@/services/api/settings/academic-period/class-schedule/get-class-schedule-template";
import { TableClassScheduleTemplate } from "./table-class-schedule-template";
import { TableClassScheduleSession } from "./table-class-schedulte-session";
import { getClassSchedule } from "@/services/api/settings/academic-period/class-schedule/get-class-schedule";

export default async function ClassSchedule({
  classId,
  isDetail,
  searchParams,
}: {
  classId: string;
  isDetail?: boolean;
  searchParams: { [key: string]: string | string[] | undefined };
}) {
  const page = Number(searchParams.page) || 1;
  const search = searchParams.q || "";
  const classScheduleTemplate = await getClassScheduleTemplate({ classId });

  const classSchedule = await getClassSchedule({
    classId,
    queryParam: {
      page,
      search: search as string,
      limit: 20,
    },
  });

  return (
    <div className="d-flex flex-column gap-3">
      <TableClassScheduleTemplate
        isDetail={isDetail}
        data={classScheduleTemplate}
        classId={classId}
      />
      <TableClassScheduleSession
        data={classSchedule}
        searchParams={searchParams}
        classId={classId}
        isDetail={isDetail}
      />
    </div>
  );
}
