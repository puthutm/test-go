"use client";

import { z } from "zod";

export const classScheduleTemplateFormSchema = z.object({
  day_name: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik harus dipilih",
    }
  ),
  start_time: z
    .array(z.date(), {
      message: "Jam mulai wajib diisi",
    })
    .min(1, "Jam mulai wajib diisi"),
  end_time: z
    .array(z.date(), {
      message: "Jam mulai wajib diisi",
    })
    .min(1, "Jam mulai wajib diisi"),
  type_of_meeting: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Periode akademik harus dipilih",
    }
  ),
});

export type ClassScheduleFormTemplateSchemaType = z.infer<
  typeof classScheduleTemplateFormSchema
>;
