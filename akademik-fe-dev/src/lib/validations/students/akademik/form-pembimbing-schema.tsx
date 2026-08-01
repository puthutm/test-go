import { z } from "zod";

const isBrowser = typeof window !== "undefined";

export const FormPembimbingSchema = z.object({
    //   Topik Konsultasi
    topik_konsultasi: z
    .string({
        required_error: "Isi topik harus diisi",
      })
      .min(1, { message: "Isi topik tidak boleh kosong" }),

    //   Isi Pesan
    isi_pesan: z
    .string({
      required_error: "Isi pesan harus diisi",
    })
    .min(1, { message: "Isi pesan tidak boleh kosong" }),

    //   Dokumen
    dokumen: isBrowser
        ? z.instanceof(FileList, { message: "Dokumen harus berupa file" }).optional()
        : z.any().optional(), // Hindari error saat SSR

});

export type FormPembimbingSchemaType = z.infer<typeof FormPembimbingSchema>;
