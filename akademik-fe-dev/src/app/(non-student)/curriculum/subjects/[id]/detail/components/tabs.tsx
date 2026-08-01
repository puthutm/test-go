"use client";
import React from "react";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useSearchParams } from "next/navigation";
import { createSearchParams } from "@/lib/utils/create-search-params";

interface ITabs {
  id: string;
  title: string;
  query: string;
}
function TabsSectionLecture() {
  const { id } = useParams();
  const searchParams = useSearchParams();

  const listTabs: ITabs[] = [
    {
      id: "1",
      title: "Data Mata Kuliah",
      query: "subject",
    },
    {
      id: "2",
      title: "CPL dan CPMK",
      query: "cpl",
    },
    {
      id: "3",
      title: "Detail RPS",
      query: "rps",
    },
    {
      id: "4",
      title: "Rencana Pembelajaran",
      query: "study",
    },
    {
      id: "5",
      title: "Rencana Evaluasi",
      query: "evaluation",
    },
  ];

  return (
    <section className="position-relative d-flex flex-wrap">
      {listTabs.map((tab: ITabs) => {
        return (
          <Link
            key={tab.id}
            className={`btn-tabs-lecturer-subject  fw-medium px-4
                    ${
                      tab.query === searchParams.get("tab")
                        ? "btn-tabs-lecturer-subject-active"
                        : "btn-tabs-lecturer-subject-not-active"
                    }    
                        `}
            color="transparent"
            href={`/curriculum/subjects/${id}/detail?${createSearchParams(
              "tab",
              tab.query
            )}`}
          >
            {tab.title}
          </Link>
        );
      })}
    </section>
  );
}

export default TabsSectionLecture;
