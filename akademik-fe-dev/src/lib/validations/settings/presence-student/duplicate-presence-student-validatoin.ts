import { z } from "zod";

export const formDuplicatePresenceSchema = z.object({
  academic_periode_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik harus dipilih",
    }
  ),
  academic_periode_id_target: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik harus dipilih",
    }
  ),
  study_program_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Program studi harus dipilih",
    }
  ),
});

export type FormDuplicatePresenceSchemaType = z.infer<
  typeof formDuplicatePresenceSchema
>;
