'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getCourseContract = async (idClass:string) : Promise<ApiResponse<ICourseContract>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data kontrak kuliah')
    }
}