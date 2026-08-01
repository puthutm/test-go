import React from 'react'

import type { Metadata } from 'next'
 
export async function generateMetadata(
): Promise<Metadata> {
  return {
    title: 'Mata Kuliah Kordinator',
  }
}
function LayoutSubjectCordination({children}:{children:React.ReactNode}) {
  return children
}

export default LayoutSubjectCordination