import { z } from "zod";

export const CreditLimitFormSchema = z.object({
  ips_min:  z
    .string()
    .min(1, "minimal ips harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1 && Number(val) <= 4,
      "ips  minimal harus berupa lebih dari 1 dan dibawah 4"
    ),
  ips_max:   z
    .string()
    .min(1, "maksimal ips harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1  && Number(val) <= 4,
      "ips  maksimal harus lebih dari 1 dan dibawah 4"
    ),
  sks_limit: z.string().min(1,{ message: "batas sks harus diisi" }),
});

export type CreditLimitFormType = z.infer<typeof CreditLimitFormSchema>;

