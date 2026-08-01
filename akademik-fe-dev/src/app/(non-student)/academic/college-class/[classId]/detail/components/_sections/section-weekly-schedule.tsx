"use client";
import React from "react";
import { useParams } from "next/navigation";
import DataTables from "@/components/ui/datatable";
import useColumnWeeklySchedule from "../_columns/column-definition-weekly-schedule";

import { useGetWeeklySchedule } from "@/services/api/academic/lecturer/class-schedule/detail-class/weekly-schedule/use-get-weekly-schedule";
function SectionWeeklySchedule() {
  const params = useParams();
  const { data: dataWeeklySchedule, isLoading: isLoadingWeeklySchedule } =
    useGetWeeklySchedule(params.classId as string);

  const { columns } = useColumnWeeklySchedule();
  return (
    <section className="position-relative table-responsive">
      <DataTables
        columns={columns}
        data={{
          data: dataWeeklySchedule?.data,
        }}
        isPaginate={false}
        pageCount={0}
        pagination={0}
        setPagination={() => {
          return;
        }}
        isLoading={isLoadingWeeklySchedule}
        total={0}
      />
    </section>
  );
}

export default SectionWeeklySchedule;
