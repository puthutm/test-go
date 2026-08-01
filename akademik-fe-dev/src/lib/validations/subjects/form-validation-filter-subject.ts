import { z} from 'zod'


const schemeFormFilterSubject = z.object({
    year_curriculum:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    subject_type:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    subject_group:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
})

export default schemeFormFilterSubject
export type IFormSchemeFilterSubject = z.infer< typeof schemeFormFilterSubject> 