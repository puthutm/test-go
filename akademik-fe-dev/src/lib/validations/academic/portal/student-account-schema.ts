import { z } from "zod";
import { NIK_REGEX } from "@/lib/constants/nik-regex";

export const StudentAccountSchema = z.object({
  batch_detail_id: z
    .string({ required_error: "Batch detail harus diisi" })
    .uuid({ message: "ID detail batch tidak valid" }),
  nik: z
    .string({ required_error: "NIK harus diisi" })
    .min(16, { message: "NIK harus 16 angka" })
    .max(16, { message: "NIK harus 16 angka" })
    .regex(NIK_REGEX, { message: "NIK tidak valid" }),
  name: z
    .string({ required_error: "Nama lengkap harus diisi" })
    .min(1, { message: "Nama lengkap harus diisi" }),
  email: z
    .string({ required_error: "Email harus diisi" })
    .email({ message: "Format email tidak valid" }),
  password: z
    .string()
    .min(8, { message: "Password minimal 8 karakter" })
    .optional()
    .or(z.literal(""))
    .default("12345678"),
  phone: z
    .string({ required_error: "Nomor telepon harus diisi" })
    .min(10, { message: "Nomor telepon minimal 10 angka" })
    .max(15, { message: "Nomor telepon maksimal 15 angka" })
    .regex(/^[0-9]+$/, { message: "Nomor telepon hanya boleh berisi angka" }),
});

export type StudentAccountSchemaType = z.infer<typeof StudentAccountSchema>;
