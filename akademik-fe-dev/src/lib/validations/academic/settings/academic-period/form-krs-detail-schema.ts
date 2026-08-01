import { z } from "zod";

export const FormKrsDetailSchema = z.object({
    class_id: z.string().min(1, "ID Kelas harus diisi")
})

export type FormKrsDetailType = z.infer<typeof FormKrsDetailSchema>