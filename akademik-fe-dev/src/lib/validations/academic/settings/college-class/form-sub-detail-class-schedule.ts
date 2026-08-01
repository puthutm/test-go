import { z } from "zod";

const MAX_FILE_SIZE = 10 * 1024 * 1024;
const ACCEPTED_FILE_TYPES = ["application/pdf"];

export const FormSubDetailClassSchedule = z.object({
  material_attachment_file: z
    .union([
      z.any().refine((file) => file instanceof File, {
      message: "Lampiran materi harus berupa file.",
      })
      .refine((file) => file?.size <= MAX_FILE_SIZE, {
        message: "Ukuran maksimal Lampiran adalah 10MB.",
      })
      .refine((file) => ACCEPTED_FILE_TYPES.includes(file?.type), {
        message: "File harus .pdf",
      }),
      z.string()
    ]),
  attendance_document_file: z.union([
    z
    .any()
    .refine((file) => file instanceof File, {
      message: "Dokumen presensi harus berupa file.",
    })
    .refine((file) => file?.size <= MAX_FILE_SIZE, {
      message: "Ukuran maksimal Dokumen adalah 10MB.",
    })
    .refine((file) => ACCEPTED_FILE_TYPES.includes(file?.type), {
      message: "File harus .pdf",
    }),
    z.string()
  ]),
  journal_document_file: z.union([
    z
    .any()
    .refine((file) => file instanceof File, {
      message: "Dokumen jurnal harus berupa file.",
    })
    .refine((file) => file?.size <= MAX_FILE_SIZE, {
      message: "Ukuran maksimal Dokumen adalah 10MB.",
    })
    .refine((file) => ACCEPTED_FILE_TYPES.includes(file?.type), {
      message: "File harus .pdf",
    }),
    z.string()
  ]),
    material_plan:z.string().min(1,{message:"Rencana materi harus diisi."}),
    material_realization:z.string().min(1,{message:"Rencana materi harus diisi."}),
});

export type FormSubDetailClassScheduleType = z.infer<
  typeof FormSubDetailClassSchedule
>;
