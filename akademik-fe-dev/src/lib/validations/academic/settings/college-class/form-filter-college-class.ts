import { z} from 'zod'


const schemeFormFilterCollegeClass = z.object({
    program_studi:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    sistem_kuliah:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    jenis_status:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    prodi_pengampu:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    kurikulum:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
    kelas:z.array(z.object({
        label: z.string().optional().default(''),
        value: z.string().optional().default('')
    })).optional().default([]),
})

export default schemeFormFilterCollegeClass
export type IFormSchemeFilterCollegeClass = z.infer< typeof schemeFormFilterCollegeClass> 