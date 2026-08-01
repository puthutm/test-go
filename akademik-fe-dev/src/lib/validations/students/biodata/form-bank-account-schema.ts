import { z } from "zod";

export const FormBankAccountSchema = z.object({
  bank_id: z
    .object({
      label: z.string().nullable(),
      value: z.string().nullable(),
    })
    .optional()
    .nullable(),
  account_name: z.string().optional(),
  account_number: z.string().optional(),
  account_file_path: z
    .instanceof(File)
    .optional()
    .refine(
      (file) => {
        if (!file) return true;
        return file.size <= 2 * 1024 * 1024; // 2mb max
      },
      { message: "File maksimal 2mb" }
    )
    .refine(
      (file) => {
        if (!file) return true;
        return file.type === "image/png" || file.type === "image/jpeg";
      },
      { message: "File dengan format png dan jpg" }
    ),
});

export type FormBankAccountSchemaType = z.infer<typeof FormBankAccountSchema>;
