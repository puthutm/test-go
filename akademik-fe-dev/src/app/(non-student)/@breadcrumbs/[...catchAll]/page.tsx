import Link from "next/link";
import { Col, Row } from "reactstrap";

type Props = {
  params: {
    catchAll: string[];
  };
};

const translateToId = (breadcrumb: string) => {
  switch (breadcrumb.toLocaleLowerCase()) {
    case "add":
      return "Tambah";

    case "settings":
      return "Pengaturan";
    case "academic period":
      return "Periode Akademik";
    case "credit limit":
      return "Batas SKS";
    case "curriculum year":
      return "Tahun Kurikulum";
    case "grade composition":
      return "Komposisi Nilai";
    case "grade scale":
      return "Skala Nilai";
    case "academic":
      return "Akademik";
    case "curriculum":
      return "Kurikulum";
    case "subjects":
      return "Mata Kuliah";
    case "college curriculum":
      return "Kurikulum Kuliah";
    case "subject cordination":
      return "Kordinator Mata Kuliah";
    case "college class":
      return "Kelas Kuliah";
    case "study program":
      return "Program Studi";
    case "classes":
      return "Kelas";
    case "detail class schedule":
      return "Detail Jadwal Perkuliahan";
    case "lectures":
      return "Perkuliahan";
    case "thesis proposal":
      return "Proposal Tugas Akhir";
    case "practical lecture proposal":
      return "Proposal Kuliah Praktek";
    case "academic year":
      return "Tahun Ajaran";
    case "curriculum study program":
      return "Kurikulum Prodi";
    case "subject":
      return "Mata Kuliah";
    case "presence student":
      return "Presensi Mahasiswa";
    default:
      return breadcrumb;
  }
};

export default function BreadcrumbSlot({ params: { catchAll } }: Props) {
  // change the last item to be in regular string, not array
  const fixedCatchAll = catchAll.map((item) => {
    if (Array.isArray(item)) {
      return item.join("/");
    }
    // remove the uuid from the array
    if (item.length === 36) {
      return "";
    }

    return item.replaceAll("-", " ");
  });

  return (
    <>
      <Row>
        <Col>
          <div className="page-title-box d-sm-flex align-items-center justify-content-between">
            <h4 className="mb-sm-0 fw-bolder" style={{ color: "#495057" }}>
              {translateToId(fixedCatchAll[0])}
            </h4>

            <div className="page-title-right">
              <ol className="breadcrumb m-0">
                {catchAll.map((_, index) => {
                  return translateToId(fixedCatchAll[index] ?? "") ===
                    "" ? null : (
                    <li
                      key={index}
                      className="breadcrumb-item active text-capitalize"
                    >
                      {index === 0 ? (
                        <Link href={`#`}>
                          {translateToId(fixedCatchAll[index] ?? "")}
                        </Link>
                      ) : (
                        <Link
                          href={`/${catchAll.slice(0, index + 1).join("/")}`}
                        >
                          {translateToId(fixedCatchAll[index] ?? "")}
                        </Link>
                      )}
                    </li>
                  );
                })}
              </ol>
            </div>
          </div>
        </Col>
      </Row>
    </>
  );
}
