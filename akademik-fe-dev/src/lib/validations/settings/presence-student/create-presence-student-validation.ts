import { z } from "zod";

export const formPresenceSchema = z
  .object({
    academic_periode_id: z.object(
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

    use_open_session: z.boolean().default(false),
    open_session_percentage: z.number().min(0).max(100),
    use_document_material: z.boolean().default(false),
    document_material_percentage: z.number().min(0).max(100),
    use_quiz: z.boolean().default(false),
    quiz_percentage: z.number().min(0).max(100),
    use_task: z.boolean().default(false),
    task_percentage: z.number().min(0).max(100),
    use_video: z.boolean().default(false),
    video_percentage: z.number().min(0).max(100),
    use_uts: z.boolean().default(false),
    uts_percentage: z.number().min(0).max(100),
    use_uas: z.boolean().default(false),
    uas_percentage: z.number().min(0).max(100),
    use_comment: z.boolean().default(false),
    comment_percentage: z.number().min(0).max(100),
  })
  .superRefine((data, ctx) => {
    const total =
      data.open_session_percentage +
      data.document_material_percentage +
      data.quiz_percentage +
      data.task_percentage +
      data.video_percentage +
      data.comment_percentage;

    if (total !== 100) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: `Total persentase harus sama dengan 100% (sekarang ${total}%)`,
        path: ["presence_percentage"],
      });
    }
  });

export type FormPresenceSchemaType = z.infer<typeof formPresenceSchema>;
