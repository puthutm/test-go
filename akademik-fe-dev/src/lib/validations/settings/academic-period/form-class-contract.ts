import { z } from "zod";

const MAX_FILE_SIZE = 10 * 1024 * 1024;
const ACCEPTED_FILE_TYPES = ["application/pdf"];

export const FormClassContractSchema = z.object({
  contract_description: z.any().refine(
    (val) => {
      const string = JSON.stringify(val);
      const json = JSON.parse(string);

      const children =
        json?.root?.children?.[0]?.children?.length ??
        json?.root?.children?.length;

      return children;
    },
    {
      message: "Deskripsi kontrak kuliah wajib diisi.",
    }
  ),
  contract_file: z
    .any()
    .refine((file) => file instanceof File, {
      message: "Berkas kontrak harus berupa file.",
    })
    .refine((file) => file?.size <= MAX_FILE_SIZE, {
      message: "Ukuran maksimal berkas adalah 10MB.",
    })
    .refine((file) => ACCEPTED_FILE_TYPES.includes(file?.type), {
      message: "File harus .pdf",
    }),
});

export type FormClassContractSchemaType = z.infer<
  typeof FormClassContractSchema
>;
