// sample data for table-khs.tsx
type SemesterData = {
  semester: number;
  year: number;
  period: string;
  courses: Course[];
  totalSKS: number;
  ips: number;
};

type Course = {
  no: number;
  kode: string;
  mataKuliah: string;
  dosen: string;
  sks: number;
  nilai: number | string;
  bobot: string;
};

export const semesters: SemesterData[] = [
  {
    semester: 4,
    year: 2024,
    period: "Genap",
    totalSKS: 24,
    ips: 3.75,
    courses: [
      {
        no: 1,
        kode: "KPSI502",
        mataKuliah: "Kalkulus",
        dosen: "Vika febri muliati, S.KOM, M.Kom",
        sks: 3,
        nilai: 50,
        bobot: "D",
      },
      ...Array(7)
        .fill(0)
        .map((_, i) => ({
          no: i + 2,
          kode: "Text",
          mataKuliah: "Text",
          dosen: "Text",
          sks: 3,
          nilai: "Text",
          bobot: "Text",
        })),
    ],
  },
  {
    semester: 3,
    year: 2024,
    period: "Ganjil",
    totalSKS: 24,
    ips: 3.75,
    courses: [
      {
        no: 1,
        kode: "KPSI502",
        mataKuliah: "Kalkulus",
        dosen: "Vika febri muliati, S.KOM, M.Kom",
        sks: 3,
        nilai: 80,
        bobot: "A",
      },
      ...Array(7)
        .fill(0)
        .map((_, i) => ({
          no: i + 2,
          kode: "Text",
          mataKuliah: "Text",
          dosen: "Text",
          sks: 3,
          nilai: "Text",
          bobot: "Text",
        })),
    ],
  },
  {
    semester: 2,
    year: 2023,
    period: "Genap",
    totalSKS: 24,
    ips: 3.75,
    courses: [
      {
        no: 1,
        kode: "KPSI502",
        mataKuliah: "Kalkulus",
        dosen: "Vika febri muliati, S.KOM, M.Kom",
        sks: 3,
        nilai: 80,
        bobot: "A",
      },
      ...Array(7)
        .fill(0)
        .map((_, i) => ({
          no: i + 2,
          kode: "Text",
          mataKuliah: "Text",
          dosen: "Text",
          sks: 3,
          nilai: "Text",
          bobot: "Text",
        })),
    ],
  },
  {
    semester: 1,
    year: 2023,
    period: "Ganjil",
    totalSKS: 24,
    ips: 3.75,
    courses: [
      {
        no: 1,
        kode: "KPSI502",
        mataKuliah: "Kalkulus",
        dosen: "Vika febri muliati, S.KOM, M.Kom",
        sks: 3,
        nilai: 80,
        bobot: "A",
      },
      ...Array(7)
        .fill(0)
        .map((_, i) => ({
          no: i + 2,
          kode: "Text",
          mataKuliah: "Text",
          dosen: "Text",
          sks: 3,
          nilai: "Text",
          bobot: "Text",
        })),
    ],
  },
];
