import { z } from "zod";

export const formKrsStatusSchema = z
  .object({
    item_status: z.enum(["rejected", "approved"]),
    reject_reason: z.string().optional(),
  })
  .superRefine((data, ctx) => {
    if (data.item_status === "rejected" && !data.reject_reason?.trim()) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Alasan penolakan wajib diisi",
        path: ["reject_reason"],
      });
    }
  });

export type FormKrsStatusSchemaType = z.infer<typeof formKrsStatusSchema>;
