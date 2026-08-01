'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getClassParticipantClassSchedule = async (idClass:string) : Promise<ApiResponse<IClassParticipant[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/class-participants`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data peserta kelas')
    }
}