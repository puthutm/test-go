import React from 'react'

import type { Metadata } from 'next'
 
export async function generateMetadata(
): Promise<Metadata> {
  return {
    title: 'Kelas Kuliah',
  }
}
function LayoutCollegeClass({children}:{children:React.ReactNode}) {
  return children
}

export default LayoutCollegeClass