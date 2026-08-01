import { z } from "zod";

export const DuplicateGradeCompositionSchema = z.object({
  academicPeriodIdSource: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik asal harus dipilih",
    }
  ),
  academicPeriodIdTarget: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik tujuan harus dipilih",
    }
  ),
  isOverWrite: z.boolean().optional().default(false),
});

export type DuplicateGradeCompositionFormType = z.infer<
  typeof DuplicateGradeCompositionSchema
>;
