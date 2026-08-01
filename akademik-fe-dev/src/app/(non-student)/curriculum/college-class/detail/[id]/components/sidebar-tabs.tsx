'use client'
// import Link from 'next/link'
import {  useRouter } from 'next/navigation'
import { useSearchParams } from 'next/navigation'
import { createSearchParams } from '@/lib/utils/create-search-params'

import React from 'react'
import { Button } from 'reactstrap'

interface ITabs {
    id:string,
    title:string,
    query:string
}
 const listTabs : ITabs[] =[
        {
            id:'1',
            title:"Detail Kelas",
            query:'detail-college-class'
        },
        {
            id:'2',
            title:"Dosen Pengajar",
            query:'lecturer-teaching'
        },
        {
            id:'3',
            title:"Peserta Kelas",
            query:'class-participants'
        },
        {
            id:'4',
            title:"Kontrak Kuliah",
            query:'contract-college'
        },
        {
            id:'6',
            title:"Jadwal Perkuliahan",
            query:'schedule-college'
        },
        {
            id:'7',
            title:"Presensi Kelas",
            query:'presence-class'
        },
        {
            id:'8',
            title:"Jadwal Ujian",
            query:'schedule-exam'
        },
        {
            id:'9',
            title:"Nilai Perkuliahan",
            query:'value-college'
        },
        {
            id:'10',
            title:"Rekap Kuesioner",
            query:'questionnaire-recap'
        },
        {
            id:'11',
            title:"RPS",
            query:'RPS'
        },
        {
            id:'12',
            title:"Tugas Kuliah",
            query:'assignment-college'
        },
        
    ]


function SidebarTabs() {
  const router = useRouter()
//   const params = useParams()
  const searchParams = useSearchParams()
  return (
   <section className="position-relative w-100 bg-white m-0 p-3 rounded-1 d-flex flex-column ">
    {
      listTabs.map((item)=>{
        return (
        <Button 
        disabled={item.id !== "2" && item.id !== "3" ? true : false}
        onClick={()=>router.push(`?${createSearchParams('tab',item.query)}`)}
        key={item.id} className='btn border-0 py-1.5 text-start'
        style={{
          background:searchParams.get('tab') === item.query ? '#10487A' : 'transparent',
          color:searchParams.get('tab') === item.query ? '#fff' : '#000',
        }}>
          {
            item.title
          }
        </Button>
        )
      })
    }

   </section> 
  )
}

export default SidebarTabs