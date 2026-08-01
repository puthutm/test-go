"use client";

import { useRouter, useSearchParams } from "next/navigation";

export const Tabs = ({ param }: { param: string }) => {
  const params = useSearchParams();
  const searchParams = params.get("tabs");
  const router = useRouter();
  const tabs = [
    { name: "Data Mata Kuliah", id: "subject-data" },
    { name: "CPL dan CPMK", id: "cpl-cpmk" },
    { name: "Detail RPS", id: "rps" },
    { name: "Rencana Pembelajaran", id: "course-plan" },
    { name: "Rencana Evaluasi", id: "evaluation-plan" },
  ];

  const handleTabClick = (tab: { name: string; id: string }) => {
    const params = new URLSearchParams(window.location.search);
    params.set("tabs", tab.id);
    router.replace(`/settings/subject/${param}/edit?${params.toString()}`);
  };

  return (
    <div className="d-flex gap-3 flex-wrap pb-3">
      {tabs?.map((tab, index) => (
        <button
          key={tab.id}
          className={`border-0  rounded-top text-center py-2 px-4 fw-semibold ${
            tab.id === (searchParams || "subject-data")
              ? "bg-primary text-white"
              : "bg-transparent"
          }`}
          onClick={() => handleTabClick(tab)}
          style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
          disabled={index !== 0}
        >
          {tab.name}
        </button>
      ))}
    </div>
  );
};
