"use client";

import { Nav, NavItem, NavLink } from "reactstrap";

import { classMerge } from "@/lib/utils/class-merge";
import { usePathname, useSearchParams, useRouter } from "next/navigation";
import { createSearchParams } from "@/lib/utils/create-search-params";

const tabList = [
  {
    label: "Biodata",
    key: "biodata",
  },
  {
    label: "Akademik",
    key: "academic",
  },
  {
    label: "Tingkat Akhir",
    key: "final-level",
  },
  {
    label: "Kerja Praktik",
    key: "kp",
  },
  {
    label: "Administrasi",
    key: "administration",
  },
  {
    label: "Kalender Akademik",
    key: "calendar",
  },
  {
    label: "KTM",
    key: "ktm",
  },
];

export const Tabs = () => {
  const tabParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const handleTabClick = (tabName: string) => {
    router.push(`${pathname}?${createSearchParams("tabs", tabName)}`);
  };
  return (
    <div className="d-flex justify-content-between">
      <Nav
        pills
        className="animation-nav profile-nav gap-2 gap-lg-3 flex-grow-1"
        role="tablist"
      >
        {tabList.map((tab) => (
          <NavItem
            className="fs-14"
            key={tab.key}
            onClick={() => handleTabClick(tab.key)}
          >
            <NavLink
              href="#"
              className={classMerge({
                active: tab.key === (tabParams.get("tabs") || "biodata"),
              })}
            >
              <span className="d-inline-block">{tab.label}</span>
            </NavLink>
          </NavItem>
        ))}
      </Nav>
    </div>
  );
};
