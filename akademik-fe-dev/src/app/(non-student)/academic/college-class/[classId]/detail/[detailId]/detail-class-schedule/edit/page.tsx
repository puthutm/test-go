import { Metadata } from "next";
import React from "react";
import ClientFormEditSubClassSchedule from "./components/client";

import { getDetailClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/get-detail-class-scedule-sub-detail";
export const metadata: Metadata = {
  title: "Edit Class Schedule",
};

export default async function FormEditSubClassSchedulePage({
  params,
}: {
  params: {
    classId: string;
    detailId: string;
  };
}) {
  //   const page = Number(searchParams.page) || 1;
  //   const search = searchParams.search || "";
  //   const limit = searchParams.limit || 10;

  const data = await getDetailClassScheduleSubDetail(
    params.classId,
    params.detailId
  );

  return <ClientFormEditSubClassSchedule dataDetailClassSchedule={data} />;
}
