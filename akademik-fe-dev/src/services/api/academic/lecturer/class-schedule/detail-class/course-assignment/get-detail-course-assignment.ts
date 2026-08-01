'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getDetailCourseAssignment = async (idClass:string,idCourseAssigment:string|undefined | null) : Promise<ApiResponse<ICourseAssignment>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/course-assisments/${idCourseAssigment}`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get detail data tugas kuliah')
    }
}