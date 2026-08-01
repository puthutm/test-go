import { z } from "zod";

export const FormClassAcademicPeriodDetailSchema = z
  .object({
    code: z
      .string({ message: "Kode kelas harus diisi" })
      .min(1, { message: "Kode kelas harus diisi" }),
    name: z
      .string({ message: "Nama kelas harus diisi" })
      .min(1, { message: "Nama kelas harus diisi" }),
    capacity: z
      .number({ message: "Kapasitas kelas harus diisi" })
      .min(1, { message: "Kapasitas kelas harus diisi" }),
    number_of_meeting: z
      .number({ message: "Jumlah pertemuan kelas harus diisi" })
      .min(1, { message: "Jumlah pertemuan kelas harus diisi" })
      .max(20, {
        message: "Jumlah pertemuan kelas maksimal 20",
      }),
    academic_periode_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Periode akademik harus dipilih",
      }
    ),
    program_study_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Program studi harus dipilih",
      }
    ),
    curriculum_year_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Tahun kurikulum harus dipilih",
      }
    ),
    subject_id: z
      .object(
        {
          label: z.string(),
          value: z.string(),
        },
        {
          message: "Mata kuliah harus dipilih",
        }
      )
      .nullable(),
  })
  .superRefine((data, ctx) => {
    const hasCurriculum = !!data.curriculum_year_id;
    const hasProgramStudy = !!data.program_study_id;
    const hasSubject = !!data.subject_id;
    if (hasCurriculum && hasProgramStudy && !hasSubject) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        path: ["subject_id"],
        message: "Mata kuliah harus dipilih",
      });
    }
  });

export type FormClassAcademicPeriodDetailType = z.infer<
  typeof FormClassAcademicPeriodDetailSchema
>;
