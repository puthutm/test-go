'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getWeeklySchedule = async (idClass:string) : Promise<ApiResponse<WeeklySchedule[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/weekly-schedule`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data jadwal kelas')
    }
}