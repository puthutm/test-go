import { notFound } from "next/navigation";
import React from "react";
import { Col, Row } from "reactstrap";

import { getCurriculumYearById } from "@/services/api/data-referensi/curriculum-year/get-curriculum-year-by-id";
import { formatDate } from "@/lib/utils/format-date";

export default async function CurriculumYearInfo({
  params,
}: {
  params: Promise<{ curriculumYearId: string }>;
}) {
  const { curriculumYearId } = await params;
  const data = await getCurriculumYearById(curriculumYearId);

  if (!data.data) return notFound();

  const datas = [
    {
      title: "Mulai Berlaku",
      value: data.data.academic_periode_name,
    },
    {
      title: "Tanggal Mulai",
      value: formatDate(data.data.start_date),
    },
    {
      title: "Tanggal Selesai",
      value: formatDate(data.data.end_date),
    },
    {
      title: "Keterangan",
      value: data?.data?.description || "-",
    },
  ];
  return (
    <>
      <h1 className="fs-5 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
        Tahun Kurikulum {data.data.years}
      </h1>
      <Row className="rounded mx-1 p-3" style={{ backgroundColor: "#F3F3F9" }}>
        {datas.map((item, index) => (
          <Col sm={6} lg={3} className="py-2" key={index}>
            <p style={{ color: "#909090" }}>{item.title}</p>
            <p className="mb-0">{item.value}</p>
          </Col>
        ))}
      </Row>
    </>
  );
}
