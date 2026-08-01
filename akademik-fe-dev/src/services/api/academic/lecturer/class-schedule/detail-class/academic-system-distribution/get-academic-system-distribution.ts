'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getAcademicSystemDistribution = async (idClass:string) : Promise<ApiResponse<AcademicSystemDistribution[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/academic-system-distributions`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data sistem kuliah')
    }
}