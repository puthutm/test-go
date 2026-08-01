import { z } from "zod";

// Define the schema for a single parent (either father or mother)
export const FormMotherSchema = z.object({
  // Basic information
  nama: z.string().optional(),
  email: z.string().email("Email tidak valid").optional(),
  nik: z
    .string()
    .min(16, "NIK harus 16 digit")
    .max(16, "NIK harus 16 digit")
    .optional(),
  no_hp: z
    .string()
    .regex(/^(\+62|62|0)8[1-9][0-9]{6,9}$/, "Nomor HP tidak valid")
    .optional(),
  no_telephone: z
    .string()
    .regex(/^(\+62|62|0)8[1-9][0-9]{6,9}$/, "Nomor Telephone tidak valid")
    .optional(),
  // Education and location
  pendidikan: z
    .enum(["SD", "SMP", "SMA", "D1", "D2", "D3", "D4", "S1", "S2", "S3"])
    .optional(),
  tempat_lahir: z.string().optional(),
  tanggal_lahir: z
    .string()
    .optional()

    .refine(
      (date) => {
        if (!date) return true;
        return !isNaN(Date.parse(date));
      },
      { message: "Format tanggal tidak valid" }
    ),

  // Address
  alamat: z.string().max(500, "Alamat maksimal 500 karakter").optional(),
  // Occupation and income
  pekerjaan: z
    .enum([
      "PNS",
      "Swasta",
      "Wiraswasta",
      "TNI/Polri",
      "Petani",
      "Nelayan",
      "Lainnya",
    ])
    .optional(),
  penghasilan: z.string().min(1, "Penghasilan harus berupa angka").optional(),

  // .refine(
  //   (val) => {
  //     if (!val) return true;
  //     const num = Number(val.replace(/\./g, ""));
  //     return !isNaN(num);
  //   },
  // { message: "Penghasilan harus berupa angka" }
  // ),

  // Status
  status_hidup: z.enum(["Hidup", "Meninggal"]).optional().nullable(),
  status_kekerabatan: z.enum(["Kandung", "Tiri", "Angkat", "Wali"]).optional(),
});

export type FormMotherSchemaType = z.infer<typeof FormMotherSchema>;
