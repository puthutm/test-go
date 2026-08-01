'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getClassAttendance = async (idClass:string) : Promise<ApiResponse<IClassAttendance>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/class-attendances`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data presensi kelas')
    }
}