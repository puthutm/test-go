import {z} from 'zod'

export const schemeFormCourseAssigment = z.object({
    schedule_id:z.object({
        label:z.string(),
        value:z.string()
    },{
        required_error:'jadwal perkuliahan harus di isi'
    }),
    title:z.string().min(1,'judul harus di isi'),
    description:z.string().min(1,'deskripsi harus di isi'),
    is_gradeable:z.boolean().default(false),
    is_use_deadline:z.boolean().default(false),
    deadline_of_assignment_submission:z.string().min(1,'batas waktu penyerahan harus di isi'),
    time_to_open:z.string().min(1,'waktu untuk membuka harus di isi'),
    retake:z.string().min(1,'retake untuk membuka harus di isi'),
})

export type FormSchemeCourseAssigment = z.infer<typeof schemeFormCourseAssigment>