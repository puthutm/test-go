import { z } from "zod";

export const formPresenceLecturerSchema = z
  .object({
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

export const formPresenceComponentSchema = z
  .object({
    study_program_id: z
      .string()
      .nonempty({ message: "Study Program Harus diisi" }),
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
    const openSession = data?.use_open_session
      ? Number(data?.open_session_percentage)
      : 0;
    const documentMaterial = data?.use_document_material
      ? Number(data?.document_material_percentage)
      : 0;
    const quiz = data?.use_quiz ? Number(data?.quiz_percentage) : 0;
    const task = data?.use_task ? Number(data?.task_percentage) : 0;
    const video = data?.use_video ? Number(data?.video_percentage) : 0;
    const comment = data?.use_comment ? Number(data?.comment_percentage) : 0;

    const total =
      openSession + documentMaterial + quiz + task + video + comment;

    if (total !== 100) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: `Total persentase harus sama dengan 100% (sekarang ${total}%)`,
        path: ["presence_percentage"],
      });
    }
  });

export type FormPresenceLecturerSchemaType = z.infer<
  typeof formPresenceLecturerSchema
>;

export type FormPresenceComponentSchemaType = z.infer<
  typeof formPresenceComponentSchema
>;
