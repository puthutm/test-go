import { z } from "zod";

export const FormWeeklyScheduleSchema = z.object({
  day: z.object(
    {
      value: z.string(),
      label: z.string(),
    },
    {
      message: "Hari harus dipilih",
    }
  ),
  start_hour: z
    .array(z.date(), {
      message: "Jam mulai harus diisi",
    })
    .min(1, {
      message: "Jam mulai harus diisi",
    }),
  end_hour: z
    .array(z.date(), {
      message: "Jam selesai harus diisi",
    })
    .min(1, {
      message: "Jam selesai harus diisi",
    }),
  type_meeting: z.object(
    {
      value: z.string(),
      label: z.string(),
    },
    {
      message: "Jenis pertemuan harus dipilih",
    }
  ),
});

export type FormClassDetailFormType = z.infer<typeof FormWeeklyScheduleSchema>;
