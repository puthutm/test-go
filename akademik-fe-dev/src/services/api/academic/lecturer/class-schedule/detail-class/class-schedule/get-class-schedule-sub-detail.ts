'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getClassScheduleSubDetail = async (idClass:string) : Promise<ApiResponse<PaginationData<IClassScheduleSubDetail>>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/class-schedules`,{
            method:'GET',
            cache:"no-store"
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data jadwal perkuliahan')
    }
}