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
function TabsSectionDetailCollegeClass() {
  const { classId } = useParams();
  const searchParams = useSearchParams();

  const listTabs: ITabs[] = [
    {
      id: "1",
      title: "Detail Kelas",
      query: "class",
    },
    {
      id: "2",
      title: "Dosen Pengajar",
      query: "lecturer",
    },
    {
      id: "3",
      title: "Peserta Kelas",
      query: "member",
    },
    {
      id: "4",
      title: "Kontrak Kuliah",
      query: "contract",
    },
    {
      id: "5",
      title: "Jadwal Perkuliahan",
      query: "schedule-college",
    },
    {
      id: "6",
      title: "Presensi Kelas",
      query: "presence",
    },
    {
      id: "7",
      title: "Jadwal Ujian",
      query: "schedule-exam",
    },
    {
      id: "8",
      title: "Nilai Perkuliahan",
      query: "course-grades",
    },
    {
      id: "9",
      title: "Rekap Kuesioner",
      query: "questionnaire-recap",
    },
    {
      id: "10",
      title: "RPS",
      query: "rps",
    },
    {
      id: "11",
      title: "Tugas Kuliah",
      query: "course-work",
    },
  ];

  return (
    <section className="position-relative mb-3 d-flex flex-wrap">
      {listTabs.map((tab: ITabs) => {
        return (
          <Link
            key={tab.id}
            className={`btn-tabs-lecturer-subject  fw-medium
                    ${
                      tab.query === searchParams.get("tab")
                        ? "btn-tabs-lecturer-subject-active"
                        : "btn-tabs-lecturer-subject-not-active"
                    }    
                        `}
            color="transparent"
            href={`/academic/college-class/${classId}/detail?${createSearchParams(
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

export default TabsSectionDetailCollegeClass;
