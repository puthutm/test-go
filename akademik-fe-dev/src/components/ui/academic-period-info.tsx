import { notFound } from "next/navigation";

import { CheckIcon } from "@/components/icons/check";
import { CloseIcon } from "@/components/icons/close";
import { formatDate } from "@/lib/utils/format-date";
import { getAcademicPeriodById } from "@/services/api/data-referensi/academic-period/get-academic-period-by-id";
import { checkOpenCloseGradeByAcademicPeriod } from "@/services/api/settings/academic-period/open-close-grade/check-open-close-grade-by-academic-period";
import { ButtonOpenCloseGrade } from "@/app/(non-student)/settings/academic-period/[academicPeriodId]/classes/components/button-open-close-grade";

export default async function AcademicPeriodInfo({
  params,
}: {
  params: Promise<{ academicPeriodId: string }>;
}) {
  const { academicPeriodId } = await params;
  const data = await getAcademicPeriodById(academicPeriodId);
  const dataOpenCloseGrade = await checkOpenCloseGradeByAcademicPeriod({
    academicPeriodId,
  });

  if (!data.data) return notFound();

  const datas = [
    {
      title: "Kode Periode",
      value: data.data.code,
    },
    {
      title: "Tahun Ajaran",
      value: data.data.academic_year,
    },
    {
      title: "Semester",
      value: data.data.semester,
    },
    {
      title: "Nama Periode",
      value: data.data.fullname,
    },
    {
      title: "Nama Singkat",
      value: data.data.shortname,
    },
    {
      title: "Jumlah Pertemuan",
      value: data.data.number_of_lecture_meeting,
    },
    {
      title: "Tanggal Awal Kuliah",
      value: formatDate(data.data.start_date_of_college),
    },
    {
      title: "Tanggal Akhir Kuliah",
      value: formatDate(data.data.end_date_of_college),
    },
    {
      title: "Tanggal Awal UTS",
      value: formatDate(data.data.start_date_of_uts),
    },
    {
      title: "Tanggal Akhir UTS",
      value: formatDate(data.data.end_date_of_uts),
    },
    {
      title: "Tanggal Awal UAS",
      value: formatDate(data.data.start_date_of_uas),
    },
    {
      title: "Tanggal Akhir UAS",
      value: formatDate(data.data.end_date_of_uas),
    },
    {
      title: "Aktif",
      value: data.data.is_active,
    },
  ];
  return (
    <>
      <div className="d-flex justify-content-between">
        <h1 className="fs-5 mb-0 fw-medium" style={{ color: "#3A3A3A" }}>
          {data.data.fullname}
        </h1>
        <ButtonOpenCloseGrade
          academicPeriodId={academicPeriodId}
          statusLock={Boolean(dataOpenCloseGrade?.data?.status_lock)}
        />
      </div>
      <div
        className="rounded d-flex flex-column p-3"
        style={{ backgroundColor: "#F3F3F9" }}
      >
        <div className="row row-cols-5 px-4">
          {datas.map((val, index) => (
            <div className="col d-flex flex-column gap-1 p-2" key={index}>
              <p
                className="fw-medium"
                style={{ color: "#3A3A3A", fontSize: "12px" }}
              >
                {val.title}
              </p>
              <p style={{ fontSize: "14px", color: "#495057" }}>
                {val.value === true ? (
                  <CheckIcon />
                ) : val.value === false ? (
                  <CloseIcon />
                ) : (
                  val.value
                )}
              </p>
            </div>
          ))}
        </div>
      </div>
    </>
  );
}
