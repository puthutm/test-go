import { z } from "zod";

export const CurriculumYearSchema = z.object({
  years: z
    .string({ required_error: "Tahun harus diisi" })
    .min(1, { message: "Tahun harus diisi" }),
  starts: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Mulai berlaku harus diisi",
    }
  ),
  start_date: z
    .array(z.date(), {
      message: "Tanggal dimulai harus diisi",
    })
    .min(1, {
      message: "Tanggal dimulai harus diisi",
    }),
  end_date: z
    .array(z.date(), {
      message: "Tanggal selesai harus diisi",
    })
    .min(1, {
      message: "Tanggal selesai harus diisi",
    }),
  description: z.string().optional(),
});

export type CurriculumYearFormType = z.infer<typeof CurriculumYearSchema>;
