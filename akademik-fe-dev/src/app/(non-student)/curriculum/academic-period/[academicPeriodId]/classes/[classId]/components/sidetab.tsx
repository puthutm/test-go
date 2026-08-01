"use client";

import { usePathname, useSearchParams, useRouter } from "next/navigation";
import { Button } from "reactstrap";

import { classMerge } from "@/lib/utils/class-merge";

export const SideTab = () => {
  const param = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const tabsParams = param?.get("tabs");

  const tabs = [
    {
      title: "Detail Kelas",
      key: "class-detail",
    },
    {
      title: "Dosen Pengajar",
      key: "class-lecturer",
    },
    {
      title: "Peserta Kelas",
      key: "participant",
    },
    {
      title: "Kontrak Kuliah",
      key: "class-contract",
    },
    {
      title: "Jadwal Perkuliahan",
      key: "class-schedule",
    },
    {
      title: "Presensi Kelas",
      key: "attendance",
    },
    {
      title: "Jadwal Kuliah",
      key: "schedule",
    },
    {
      title: "Nilai Perkuliahan",
      key: "class-score",
    },
    {
      title: "Rekap Kuisioner",
      key: "questionnaire-summary",
    },
    {
      title: "RPS",
      key: "rps",
    },
    {
      title: "Tugas Kuliah",
      key: "homework",
    },
  ];

  const handleTabClick = (tabName: string) => {
    const params = new URLSearchParams(param.toString());

    params.set("tabs", tabName);

    router.push(`${pathname}?${params.toString()}`);
  };

  return (
    <div className="d-flex flex-lg-column gap-1 flex-wrap flex-lg-nowrap">
      {tabs.map((tab) => (
        <Button
          key={tab.key}
          color={
            tab.key === (tabsParams || tabs[0].key) ? "primary" : "transparent"
          }
          className={classMerge(
            tab.key === (tabsParams || tabs[0].key)
              ? "text-white"
              : "text-muted",
            "text-start d-flex align-items-center gap-3"
          )}
          onClick={() => handleTabClick(tab.key)}
        >
          <span>{tab.title}</span>
        </Button>
      ))}
    </div>
  );
};
