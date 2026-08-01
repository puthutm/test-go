import { z } from "zod";

export const AssessmentWeightSchema = z
  .object({
    attitude_behavior_percentage: z
      .union([z.string(), z.number()])
      .refine((val) => {
        const num = Number(val);
        return !isNaN(num) && num >= 0 && num <= 100;
      }, "Persentase sikap/perilaku harus berupa angka antara 0 sampai 100"),
    task_percentage: z.union([z.string(), z.number()]).refine((val) => {
      const num = Number(val);
      return !isNaN(num) && num >= 0 && num <= 100;
    }, "Persentase tugas harus berupa angka antara 0 sampai 100"),
    uts_percentage: z.union([z.string(), z.number()]).refine((val) => {
      const num = Number(val);
      return !isNaN(num) && num >= 0 && num <= 100;
    }, "Persentase UTS harus berupa angka antara 0 sampai 100"),
    uas_percentage: z.union([z.string(), z.number()]).refine((val) => {
      const num = Number(val);
      return !isNaN(num) && num >= 0 && num <= 100;
    }, "Persentase UAS harus berupa angka antara 0 sampai 100"),
  })
  .refine(
    (data) => {
      const total =
        Number(data.attitude_behavior_percentage) +
        Number(data.task_percentage) +
        Number(data.uts_percentage) +
        Number(data.uas_percentage);
      return total === 100;
    },
    {
      message: "total persentase harus sama dengan 100%",
      path: ["attitude_behavior_percentage"],
    }
  );

export type AssessmentWeightFormType = z.infer<typeof AssessmentWeightSchema>;
