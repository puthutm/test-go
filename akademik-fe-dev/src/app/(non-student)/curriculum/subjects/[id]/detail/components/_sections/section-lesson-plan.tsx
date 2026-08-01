"use client";
import React, { useState } from "react";

// import third party component
import Link from "next/link";
// import component
import DataTables from "@/components/ui/datatable";
import ModalDetailLessonPlan from "../modal-detail-lesson-plan";
import useColumnLessonPlan from "../_columns/column-definition-lesson-plan";

export interface IDummyValueTable {
  id: string;
  sesi: string;
  sub_cpk: string;
  penilaian: string;
  metode_pembelajaran: string;
  materi_pembelajaran: string;
  bobot: string;
}

export interface IModalDetailLessonPlan {
  status: boolean;
  title: "Detail Rencana Pembelajaran";
  data?: string;
}

const dummyValueTable: PaginationData<IDummyValueTable> = {
  metadata: {
    page: 1,
    size: 1,
    total_data: 1,
    total_page: 1,
  },
  data: [
    {
      id: "1",
      sesi: "1",
      sub_cpk: "Pengantar Interaksi Manusia dan Kompute",
      penilaian: "Pengantar Interaksi Manusia dan Komputer",
      metode_pembelajaran: "text",
      materi_pembelajaran: " Pengantar Interaksi Manusia dan Komputer ",
      bobot: "text",
    },
  ],
};

function SectionLessonPlan() {
  const [queryParams] = useState<QueryParam>({
    page: 1,
  });

  //! state show MODAL detail
  const [showModalDetailLesson, setShowModalDetailLesson] =
    useState<IModalDetailLessonPlan>({
      status: false,
      title: "Detail Rencana Pembelajaran",
    });

  const { columns } = useColumnLessonPlan(setShowModalDetailLesson);

  return (
    <>
      <ModalDetailLessonPlan
        showModal={showModalDetailLesson}
        setShowModal={setShowModalDetailLesson}
      />
      {/*//! tables cpl */}
      <section className="mt-3 position-relative table-responsive">
        <DataTables
          columns={columns}
          data={dummyValueTable}
          pageCount={dummyValueTable?.metadata.total_page as number}
          pagination={queryParams}
          setPagination={() => {}}
          isLoading={false}
          total={dummyValueTable.metadata.total_data as number}
        />
      </section>
      {/*//! section back */}
      <section className="position-relative mt-3 d-flex justify-content-starts">
        <Link
          href={"?tab=rps"}
          className="btn  d-flex align-items-center gap-2"
          color="transparent"
          style={{ color: "#10487A", border: "1px solid #10487A" }}
        >
          Kembali
        </Link>
      </section>
    </>
  );
}

export default SectionLessonPlan;
