'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getDetailClassScheduleSubDetail = async (idClass:string,idClassSchedule:string) : Promise<ApiResponse<IClassScheduleSubDetail>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/class-schedules/${idClassSchedule}`,{
            method:'GET',
            cache:'no-store'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data detail jadwal perkuliahan')
    }
}