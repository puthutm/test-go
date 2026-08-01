import { z } from "zod";

export const FormClassForProgramHeadSchema = z
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
    // .max(50, {
    //   message: "Kapasistas kelas maksimal 50",
    // }),
    number_of_meeting: z
      .number({ message: "Jumlah pertemuan kelas harus diisi" })
      .min(1, { message: "Jumlah pertemuan kelas harus diisi" })
      .max(20, {
        message: "Jumlah pertemuan kelas maksimal 20",
      }),
    academic_period_id: z
      .string({ required_error: "Periode akademik harus dipilih" })
      .min(1, { message: "Periode akademik harus dipilih" }),
    curriculum_year_id: z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Tahun kurikulum harus dipilih",
      }
    ),
    program_study_id: z
      .object({
        label: z.string(),
        value: z.string(),
      })
      .nullable()
      .optional(),
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
    const hasSubject = !!data.subject_id;
    if (hasCurriculum && !hasSubject) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        path: ["subject_id"],
        message: "Mata kuliah harus dipilih",
      });
    }
  });

export type FormClassForProgramHeadType = z.infer<
  typeof FormClassForProgramHeadSchema
>;
