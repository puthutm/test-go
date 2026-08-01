"use client";

import { usePathname, useSearchParams, useRouter } from "next/navigation";
import { Button } from "reactstrap";

import { classMerge } from "@/lib/utils/class-merge";

export const SideTab = ({
  tabs,
}: {
  tabs: {
    label: string;
    key: string;
    icon: JSX.Element;
  }[];
}) => {
  const param = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const stateParams = param?.get("state");
  const tabsParams = param?.get("tabs");

  const sideTabList = [
    {
      tab: "biodata",
      state: "biodata",
    },
    {
      tab: "academic",
      state: "khs",
    },
    {
      tab: "final-level",
      state: "proposal",
    },
    {
      tab: "kp",
      state: "info",
    },
    {
      tab: "administration",
      state: "mail",
    },
  ];

  const active = sideTabList.find((tab) => tab.tab === tabsParams);

  const handleTabClick = (tabName: string) => {
    const params = new URLSearchParams(param.toString());

    params.get("tabs");
    params.set("state", tabName);

    router.push(`${pathname}?${params.toString()}`);
  };

  return (
    <div className="d-flex flex-lg-column gap-1 flex-wrap flex-lg-nowrap">
      {tabs.map((tab) => (
        <Button
          key={tab.key}
          color={
            tab.key === (stateParams || active?.state || "biodata")
              ? "primary"
              : "transparent"
          }
          className={classMerge(
            tab.key === (stateParams || active?.state || "biodata")
              ? "text-white"
              : "text-muted",
            "text-start d-flex align-items-center gap-3"
          )}
          onClick={() => handleTabClick(tab.key)}
        >
          {tab.icon}
          <span>{tab.label}</span>
        </Button>
      ))}
    </div>
  );
};
