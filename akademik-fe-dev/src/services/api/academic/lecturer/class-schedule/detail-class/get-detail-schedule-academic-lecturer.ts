'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getDetailClassScheduleLecturer = async (idClass:string) : Promise<ApiResponse<ClassScheduleDetail>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data kelas kuliah')
    }
}