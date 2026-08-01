import { z } from "zod";

export const GradeCompositionSchema = z.object({
  academic_periode_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode Akademik harus dipilih",
    }
  ),
  value_element_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "element harus dipilih",
    }
  ),
  percentage: z
    .string()
    .min(1, "persen harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1 && Number(val) <= 100,
      "persen harus berupa angka antara 1 sampai 100"
    ),
  is_passing_requirement: z.boolean().optional().default(false),
});

export type GradeCompositionFormType = z.infer<typeof GradeCompositionSchema>;
