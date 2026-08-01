"use client";
import React from "react";
import { usePathname, useRouter } from "next/navigation";
import { useSearchParams } from "next/navigation";

function TabsPresences() {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const tabParam = searchParams.get("tabs");
  const router = useRouter();

  const tabs = [
    {
      name: "Presensi",
      id: "presence",
    },
    {
      name: "Komponen Presensi",
      id: "component",
    },
  ];

  const handleTabClick = (tab: { name: string; id: string }) => {
    const params = new URLSearchParams(window.location.search);
    params.set("tabs", tab.id);
    router.replace(`${pathname}?${params.toString()}`);
  };

  return (
    <div className="d-flex gap-3 flex-wrap pb-2 border-bottom">
      {tabs?.map((tab) => (
        <button
          key={tab.id}
          className={`border-0 rounded-top text-center py-2 px-4 fw-semibold ${
            tab.id === (tabParam || "presence")
              ? "bg-primary text-white"
              : "bg-transparent"
          }`}
          onClick={() => handleTabClick(tab)}
          style={{ color: "#909090", borderRadius: "4px 4px 0px 0px" }}
        >
          {tab.name}
        </button>
      ))}
    </div>
  );
}

export default TabsPresences;
