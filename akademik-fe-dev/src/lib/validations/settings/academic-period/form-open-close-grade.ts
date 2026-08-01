import { z } from "zod";

export const formOpenCloseGradeSchema = z.object({
  status_locked: z.boolean(),
});

export type FormOpenCloseGradeSchemaType = z.infer<
  typeof formOpenCloseGradeSchema
>;
