'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getDistributionOfStudyProgram = async (idClass:string) : Promise<ApiResponse<DistributionOfStudyProgram[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/distribution-of-study-programs`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data jadwal kelas')
    }
}