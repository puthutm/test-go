import React from 'react'
import { Metadata } from 'next'
import ClientComponentCollegeClass from './_components/client'

export const metadata : Metadata = {
    title:'Kelas Kuliah'
}

function PageCollegeClassAcademic() {
  return <ClientComponentCollegeClass/>
}

export default PageCollegeClassAcademic