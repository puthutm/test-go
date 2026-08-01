"use client";

// import Link from "next/link";
import { useRouter } from "next/navigation";
import { useSearchParams } from "next/navigation";
import React from "react";

const ThesisProposalSidetab = () => {
  const searchParams = useSearchParams()
  const router = useRouter();
  let tabs = searchParams.get("tabs");

  const sidetabData = [
    // Proposal TA, Tugas Akhir, Bimbingan TA, Syarat Ujian, Jadwal Ujian, Riwayat Nilai Ujian, Nilai Akhir
    {
      title: "Proposal TA",
      slug: "proposal",
      icon: "mdi mdi-file-document-outline",
    },
    // {
    //   title: "Tugas Akhir",
    //   slug: "thesis",
    //   icon: "mdi mdi-file-outline",
    // },
    {
      title: "Bimbingan TA",
      slug: "consultation",
      icon: "mdi mdi-account-group-outline",
    },
    {
      title: "Syarat Ujian",
      slug: "requirements",
      icon: "mdi mdi-information-outline",
    },
    {
      title: "Jadwal Ujian",
      slug: "exam-schedule",
      icon: "mdi mdi-calendar-outline",
    },
    {
      title: "Riwayat Nilai Ujian",
      slug: "exam-history",
      icon: "mdi mdi-history",
    },
    {
      title: "Nilai Akhir",
      slug: "final-grade",
      icon: "mdi mdi-clipboard-outline",
    },
  ];

  if (!tabs || !sidetabData.find((tab) => tab.slug === tabs)) {
    tabs = sidetabData[0].slug;
  }

  return (
    <div className="d-flex gap-2 justify-content-between flex-column">
      {sidetabData.map((tab,index) => (
        // <Link
        //   href={`/lectures/thesis-proposal?tabs=${tab.slug}`}
        //   key={tab.slug}
        //   className={`btn w-100 text-start d-flex gap-2 ${
        //     tabs === tab.slug ? "btn-primary" : "btn-ghost-dark"
        //   }`}
        // >
        //   <i className={tab.icon}></i>
        //   <span>{tab.title}</span>
        // </Link>
         <button
         onClick={()=>{
            router.push(`?tabs=${tab.slug}`)
         }}
         disabled={index === 3 || index === 4 || index === 5 ? true : false}
          key={tab.slug}
          className={`btn w-100 text-start d-flex gap-2 border-0 ${
            tabs === tab.slug ? "btn-primary" : "btn-ghost-dark"
          }`}
        >
          <i className={tab.icon}></i>
          <span>{tab.title}</span>
        </button>
      ))}
    </div>
  );
};

export default ThesisProposalSidetab;
