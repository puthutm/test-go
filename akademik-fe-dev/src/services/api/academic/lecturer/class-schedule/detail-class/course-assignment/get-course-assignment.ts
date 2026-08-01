'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getCourseAssignment = async (idClass:string) : Promise<ApiResponse<ICourseAssignment[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/course-assisments`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data tugas kuliah')
    }
}