'use client'
import React from 'react'

import { useSearchParams } from 'next/navigation'

import SectionLectureTeaching from './components/section-lecture-teaching'
import SectionClassParticipants from './components/section-class-participants'

function PageDetailCollegeClass() {
    const searchParams = useSearchParams()
  return  searchParams.get('tab') === 'lecturer-teaching' ?
   <SectionLectureTeaching/> 
   : 
   searchParams.get('tab') === 'class-participants' ?
    <SectionClassParticipants/>
    :
   <p>not found</p>
}

export default PageDetailCollegeClass