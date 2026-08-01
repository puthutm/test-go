export const revalidate = 0;

import { Metadata } from "next";

import TabBiodata from "./components/biodata/biodata-tab";
import TabAcademic from "./components/akademik/academic-tab";
import TabFinalLevel from "./components/final-level/final-level-tab";
import KTMView from "./components/ktm";
import TabKp from "./components/kp/kp-tab";
import CalendarAcademicView from "./components/calendar/calendar-akademik";
import TabAdministration from "./components/administration/administration-tab";

export const metadata: Metadata = {
  title: "Mahasiswa",
};

export default function StudentPage({
  searchParams,
}: {
  searchParams: { [key: string]: string | undefined };
}) {
  const tabParam = Array.isArray(searchParams.tabs)
    ? searchParams.tabs[0]
    : searchParams.tabs;

  switch (tabParam) {
    case "academic":
      return <TabAcademic searchParams={searchParams} />;
    case "final-level":
      return <TabFinalLevel searchParams={searchParams} />;
    case "ktm":
      return <KTMView />;
    case "calendar":
      return <CalendarAcademicView />;
    case "kp":
      return <TabKp searchParams={searchParams} />;
    case "administration":
      return <TabAdministration searchParams={searchParams} />;
    default:
      return <TabBiodata searchParams={searchParams} />;
  }
}
